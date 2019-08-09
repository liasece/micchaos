package loginmodule

import (
	"command"
	"github.com/liasece/micserver/servercomm"
	"github.com/liasece/micserver/util"
	"playermodule/boxes"
	"reflect"
	// "time"
)

type TmpPlayer struct {
	Account *boxes.Account `json:"account"`
}

type HandlerClient struct {
	*LoginModule
	mappingFunc map[string]func(smsg *servercomm.SForwardFromGate)
}

func (this *HandlerClient) Init(mod *LoginModule) {
	this.LoginModule = mod
	this.mappingFunc = make(map[string]func(smsg *servercomm.SForwardFromGate))
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
		msgName := "command." + funcName[2:]
		this.mappingFunc[msgName] =
			hf.Method(i).Interface().(func(smsg *servercomm.SForwardFromGate))
	}
}

func (this *HandlerClient) OnRecvClientMsg(smsg *servercomm.SForwardFromGate) {
	this.Info("[Login.HandlerClient.OnRecvClientMsg] 收到 Client 消息 %s",
		smsg.MsgName)
	// 从消息处理映射集合找到对应的处理函数并且执行
	if f, ok := this.mappingFunc[smsg.MsgName]; ok {
		f(smsg)
	}
}

// 注册账号
func (this *HandlerClient) OnCS_Register(smsg *servercomm.SForwardFromGate) {
	msg := &command.CS_Register{}
	msg.ReadBinary(smsg.Data)
	this.Debug("玩家请求注册 %s", msg.GetJson())
	tmpuuid, err := util.NewUniqueID(101)
	if err != nil {
		this.Error("UUID构建错误 %s", err.Error())
		return
	}
	confirm := &TmpPlayer{}
	newaccount := &TmpPlayer{}
	newaccount.Account = &boxes.Account{}
	newaccount.Account.PassWordMD5 = msg.PassWowdMD5
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
			this.Info("查找玩家成功 %s:%s:%s:%s", newaccount.Account.LoginName,
				newaccount.Account.UUID, newaccount.Account.PhoneNumber,
				newaccount.Account.PassWordMD5)
		} else {
			this.Info("目标玩家已经存在了，创建账号失败，已存在玩家的UUID[%s]",
				confirm.Account.UUID)
			send := &command.SC_ResRigster{
				Code:      1,
				Message:   "目标用户名已存在",
				ConnectID: smsg.ClientConnID,
			}
			this.SendMsgToClient(smsg.FromServerID, smsg.ClientConnID, send)
		}
	}
}

// 玩家登陆
func (this *HandlerClient) OnCS_Login(smsg *servercomm.SForwardFromGate) {
	msg := &command.CS_Login{}
	msg.ReadBinary(smsg.Data)
	this.Info("command.CS_Login: %s", msg.GetJson())
	tmpplayer := &TmpPlayer{}
	err := this.mongo_userinfos.SelectOneByAccount(
		msg.LoginName, msg.PassWowdMD5, tmpplayer)
	if err != nil {
		// 登陆失败
		this.Error("登陆失败 Err[%s] ReqJson[%s]", err.Error(), msg.GetJson())
		send := &command.SC_ResLogin{
			Code:      1,
			Message:   "用户名或密码不匹配",
			ConnectID: smsg.ClientConnID,
		}
		this.SendMsgToClient(smsg.FromServerID, smsg.ClientConnID, send)
	} else {
		// 登陆成功
		this.Info("登陆成功 %s", msg.GetJson())
		send := &command.SC_ResLogin{
			Code:      0,
			Message:   "login secess",
			ConnectID: smsg.ClientConnID,
			Account:   tmpplayer.Account.GetMsg(),
		}
		this.SendServerCmdToServer(smsg.FromServerID, send)
	}
}
