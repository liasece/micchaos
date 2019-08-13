package client

import (
	"command"
	"github.com/liasece/micserver/msg"
	"reflect"
)

type CmdHandler struct {
	*Client
	mappingFunc map[string]func(msgbinary *msg.MessageBinary)
}

func (this *CmdHandler) Init(c *Client) {
	this.Client = c
	this.mappingFunc = make(map[string]func(msgbinary *msg.MessageBinary))
	// 创建消息处理消息的映射
	hf := reflect.ValueOf(this)
	hft := hf.Type()
	for i := 0; i < hf.NumMethod(); i++ {
		funcName := hft.Method(i).Name
		// 处理消息的方法名称必须符合规范： OnCS_(MsgSubName)
		if len(funcName) < 5 || funcName[:5] != "OnSC_" {
			continue
		}
		// 计算方法名对应的消息名
		msgName := "command." + funcName[2:]
		this.mappingFunc[msgName] =
			hf.Method(i).Interface().(func(msgbinary *msg.MessageBinary))
	}
}

func (this *CmdHandler) OnSC_ResAccountRigster(msgbinary *msg.MessageBinary) {
	msg := &command.SC_ResAccountRigster{}
	msg.ReadBinary(msgbinary.ProtoData)

	this.Conn.SendCmd(this.GetLoginMsg())
	if msg.Code != 0 {
		this.Error("注册账号失败 %s", msg.GetJson())
		return
	}
	this.Info("注册成功 %s", msg.GetJson())
	if msg.Account.LoginName != "" {
		this.Logger.SetLogName(msg.Account.LoginName)
	}
}

func (this *CmdHandler) OnSC_ResAccountLogin(msgbinary *msg.MessageBinary) {
	msg := &command.SC_ResAccountLogin{}
	msg.ReadBinary(msgbinary.ProtoData)

	if msg.Code != 0 {
		this.Error("登陆失败 %s", msg.GetJson())
		return
	}
	this.Info("登陆成功 %s", msg.GetJson())
	if msg.Account.LoginName != "" {
		this.Logger.SetLogName(msg.Account.LoginName)
	}
	this.Conn.SendCmd(&command.CS_EnterGame{})
}
