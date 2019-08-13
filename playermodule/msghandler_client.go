package playermodule

import (
	"command"
	"github.com/liasece/micserver/servercomm"
	"reflect"
	"time"
)

type HandlerClient struct {
	*PlayerModule

	lastCheckTime int64
	msgCount      int64
	mappingFunc   map[string]func(smsg *servercomm.SForwardFromGate)
}

func (this *HandlerClient) Init(mod *PlayerModule) {
	this.PlayerModule = mod
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
	this.Info("[HandlerClient.OnRecvClientMsg] 收到 Client 消息 %s",
		smsg.MsgName)
	this.msgCount++
	now := time.Now().UnixNano()
	if now-this.lastCheckTime > 1*1000*1000*1000 {
		this.lastCheckTime = now
		if this.msgCount != 0 {
			this.Error("本秒处理消息 %d", this.msgCount)
		}
		this.msgCount = 0
	}
	if f, ok := this.mappingFunc[smsg.MsgName]; ok {
		f(smsg)
	} else {
		this.Error("未知的消息 %s", smsg.MsgName)
	}
}

func (this *HandlerClient) OnCS_EnterGame(smsg *servercomm.SForwardFromGate) {
	msg := &command.CS_EnterGame{}
	msg.ReadBinary(smsg.Data)
	this.Info("收到 %s", msg.GetJson())
}
