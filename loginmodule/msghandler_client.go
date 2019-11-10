package loginmodule

import (
	"reflect"
	"time"

	"ccmd"
	"encoding/json"
	"github.com/liasece/micserver/servercomm"
	"github.com/liasece/micserver/session"
	"github.com/liasece/micserver/util/math"
	"github.com/liasece/micserver/util/monitor"
	"github.com/liasece/micserver/util/uid"
	"playermodule/boxes"
)

type TmpPlayer struct {
	Account *boxes.Account `json:"account"`
}

type HandlerClient struct {
	*LoginModule
	mappingFunc map[string]func(session session.Session, data []byte)

	// 模块的负载
	ClientMsgLoad          monitor.Load
	lastCheckClientMsgLoad int64
}

func (this *HandlerClient) Init(mod *LoginModule) {
	this.LoginModule = mod
	this.mappingFunc = make(map[string]func(session session.Session, data []byte))
	// 创建消息处理消息的映射
	hf := reflect.ValueOf(this)
	hft := hf.Type()
	for i := 0; i < hf.NumMethod(); i++ {
		funcName := hft.Method(i).Name
		// 处理消息的方法名称必须符合规范： OnCS_(MsgSubName)
		if len(funcName) < 5 || funcName[:5] != "OnCS_" {
			continue
		}
		// 计算方法名对应的消息名
		msgName := "ccmd." + funcName[2:]
		this.mappingFunc[msgName] =
			hf.Method(i).Interface().(func(session session.Session, data []byte))
	}

	this.TimerManager.RegTimer(time.Second*1, 0, false, this.watchClientMsgLoadToLog)
}

func (this *HandlerClient) OnClientMessage(smsg *servercomm.ClientMessage) {
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(smsg.Data, top)

	this.Info("[Login.HandlerClient.OnRecvClientMsg] 收到 Client 消息 "+
		"MsgName[%s] Data[%s]",
		top.MsgName, string(smsg.Data))

	this.ClientMsgLoad.AddLoad(1)

	se := session.Session{}
	se.FromMap(smsg.Session)

	// 从消息处理映射集合找到对应的处理函数并且执行
	if f, ok := this.mappingFunc[top.MsgName]; ok {
		f(se, top.Data)
	}
}

// 注册账号
func (this *HandlerClient) OnCS_AccountRegister(
	session session.Session, data []byte) {
	msg := &ccmd.CS_AccountRegister{}
	json.Unmarshal(data, msg)
	this.Debug("玩家请求注册 %s", string(data))
	tmpuuid, err := uid.NewUniqueID(101)
	if err != nil {
		this.Error("UUID构建错误 %s", err.Error())
		return
	}
	salt, errGSalt := math.GenerateRandomString(16)
	if errGSalt != nil {
		this.Error("计算生成盐错误 %s", errGSalt.Error())
		return
	}
	pswmd5ws := math.HmacSha256ByString(msg.PassWordMD5, salt)
	confirm := &TmpPlayer{}
	newaccount := &TmpPlayer{}
	newaccount.Account = &boxes.Account{}
	newaccount.Account.PassWordMD5WSSalt = salt
	newaccount.Account.PassWordMD5WS = pswmd5ws
	newaccount.Account.LoginName = msg.LoginName
	newaccount.Account.UUID = tmpuuid
	err = this.mongo_userinfos.FindOneOrCreate(msg.LoginName, newaccount, confirm)
	if err != nil {
		this.Info("数据库查询错误 Err[%s]", err.Error())
	} else {
		if confirm.Account == nil {
			this.Info("创建账号失败 confirm.Account == nil")
			return
		}
		if confirm.Account.UUID == newaccount.Account.UUID {
			this.Info("查找玩家成功 %s:%s:%s:%s:%s", newaccount.Account.LoginName,
				newaccount.Account.UUID, newaccount.Account.PhoneNumber,
				newaccount.Account.PassWordMD5WS,
				newaccount.Account.PassWordMD5WSSalt)
			send := &ccmd.SC_ResAccountRigster{
				Code:      0,
				Message:   "注册成功",
				ConnectID: session.GetConnectID(),
				Account:   newaccount.Account.GetMsg(),
			}
			this.SendMsgToClient(session.GetBind("gate"),
				session.GetConnectID(), send)
		} else {
			this.Info("目标玩家已经存在了，创建账号失败，已存在玩家的UUID[%s]",
				confirm.Account.UUID)
			send := &ccmd.SC_ResAccountRigster{
				Code:      1,
				Message:   "目标用户名已存在",
				ConnectID: session.GetConnectID(),
			}
			this.SendMsgToClient(session.GetBind("gate"),
				session.GetConnectID(), send)
			send1 := &ccmd.SC_Tips{
				Message: "换个用户名再试",
			}
			this.SendMsgToClient(session.GetBind("gate"),
				session.GetConnectID(), send1)
		}
	}
}

// 玩家登陆
func (this *HandlerClient) OnCS_AccountLogin(
	session session.Session, data []byte) {
	msg := &ccmd.CS_AccountLogin{}
	json.Unmarshal(data, msg)
	tmpplayer := &TmpPlayer{}
	err := this.mongo_userinfos.SelectOneByAccount(
		msg.LoginName, tmpplayer)
	if err != nil {
		// 登陆失败
		this.Error("登陆失败 Err[%s] ReqJson[%s]", err.Error(), string(data))
		send := &ccmd.SC_ResAccountLogin{
			Code:      1,
			Message:   "目标账号不存在",
			ConnectID: session.GetConnectID(),
		}
		this.SendMsgToClient(session.GetBind("gate"),
			session.GetConnectID(), send)
	} else {
		pswmd5ws := math.HmacSha256ByString(msg.PassWordMD5,
			tmpplayer.Account.PassWordMD5WSSalt)
		if tmpplayer.Account.PassWordMD5WS != pswmd5ws {
			// 密码错误
			this.Info("登陆失败 密码错误 ReqJson[%s]", string(data))
			send := &ccmd.SC_ResAccountLogin{
				Code:      1,
				Message:   "密码错误",
				ConnectID: session.GetConnectID(),
			}
			this.SendMsgToClient(session.GetBind("gate"),
				session.GetConnectID(), send)
		} else {
			// 登陆成功
			this.Info("登陆成功 Msg[%s] %s:%s:%s:%s:%s", string(data),
				tmpplayer.Account.LoginName, tmpplayer.Account.UUID,
				tmpplayer.Account.PhoneNumber,
				tmpplayer.Account.PassWordMD5WS,
				tmpplayer.Account.PassWordMD5WSSalt)
			session.SetUUID(tmpplayer.Account.UUID)
			session.SyncToModule(&this.BaseModule, session.GetBind("gate"))
			send := &ccmd.SC_ResAccountLogin{
				Code:      0,
				Message:   "login secess",
				ConnectID: session.GetConnectID(),
				Account:   tmpplayer.Account.GetMsg(),
			}
			this.SendMsgToClient(session.GetBind("gate"),
				session.GetConnectID(), send)
		}
	}
}

func (this *HandlerClient) SendMsgToClient(gateid string,
	to string, msg interface{}) {
	btop := ccmd.GetSCTopLayer(msg)
	this.LoginModule.SendBytesToClient(gateid, to, 0, btop)
}

func (this *HandlerClient) watchClientMsgLoadToLog(dt time.Duration) bool {
	load := this.ClientMsgLoad.GetLoad()
	incValue := load - this.lastCheckClientMsgLoad
	if incValue > 0 {
		this.Info("[HandlerClient.watchClientMsgLoadToLog] Within %d sec load:[%d]",
			int64(dt.Seconds()), incValue)
	}
	this.lastCheckClientMsgLoad = load
	return true
}
