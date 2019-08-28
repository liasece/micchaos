package client

import (
	"ccmd"
	"encoding/json"
	"reflect"
)

type CmdHandler struct {
	*Client
	mappingFunc map[string]func(data []byte)
}

func (this *CmdHandler) Init(c *Client) {
	this.Client = c
	this.mappingFunc = make(map[string]func(data []byte))
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
		msgName := "ccmd." + funcName[2:]
		this.mappingFunc[msgName] =
			hf.Method(i).Interface().(func(data []byte))
	}
}

// 注册账号的回复信息
func (this *CmdHandler) OnSC_ResAccountRigster(data []byte) {
	msg := &ccmd.SC_ResAccountRigster{}
	json.Unmarshal(data, msg)

	if msg.Code != 0 {
		this.Error("注册账号失败 %s", string(data))
	} else {
		this.Info("注册成功 %s", string(data))
		if msg.Account.LoginName != "" {
			this.Logger.SetLogName(msg.Account.LoginName)
		}
	}
	// 登陆
	this.SendMsg(this.GetLoginMsg())
}

// 登陆账号的回复信息
func (this *CmdHandler) OnSC_ResAccountLogin(data []byte) {
	msg := &ccmd.SC_ResAccountLogin{}
	json.Unmarshal(data, msg)

	if msg.Code != 0 {
		this.Error("登陆失败 %s", string(data))
		return
	}
	if msg.Account.LoginName != "" {
		this.Logger.SetLogName(msg.Account.LoginName)
	}
	this.Info("登陆成功 %s", string(data))
	// 进入游戏
	this.SendMsg(&ccmd.CS_EnterGame{})
}
