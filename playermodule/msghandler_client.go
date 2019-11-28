package playermodule

import (
	"ccmd"
	"encoding/json"
	"github.com/liasece/micserver/servercomm"
	"github.com/liasece/micserver/session"
	"github.com/liasece/micserver/util/monitor"
	"reflect"
	"time"
)

type HandlerClient struct {
	*PlayerModule

	// 模块的负载
	ClientMsgLoad          monitor.Load
	lastCheckClientMsgLoad int64

	mappingFunc map[string]func(session *session.Session, data []byte)
}

func (this *HandlerClient) Init(mod *PlayerModule) {
	this.PlayerModule = mod
	this.mappingFunc = make(map[string]func(session *session.Session, data []byte))
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
			hf.Method(i).Interface().(func(session *session.Session, data []byte))
	}

	this.TimerManager.RegTimer(time.Second*1, 0, false, this.watchClientMsgLoadToLog)
}

//
func (this *HandlerClient) OnClientMessage(session *session.Session,
	smsg *servercomm.ClientMessage) {
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(smsg.Data, top)
	this.Info("[HandlerClient.OnRecvClientMsg] 收到 Client 消息 %s",
		top.MsgName)

	this.ClientMsgLoad.AddLoad(1)

	// 根据消息名映射消息处理函数
	if f, ok := this.mappingFunc[top.MsgName]; ok {
		f(session, top.Data)
	} else {
		this.Error("未知的消息 %s", top.MsgName)
	}
}

// 客户端请求进入游戏
func (this *HandlerClient) OnCS_EnterGame(session *session.Session, data []byte) {
	msg := &ccmd.CS_EnterGame{}
	json.Unmarshal(data, msg)
	this.Info("收到 %s", string(data))
	player := this.PlayerDocManager.GetPlayerDocMust(session.GetUUID())
	if player != nil {
		player.AfterOnline(session)
	} else {
		this.Error("获取Player失败 %+v", session)
	}
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
