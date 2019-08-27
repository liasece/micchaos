package command

import (
	"encoding/json"
)

type IMsg interface {
	GetMsgName() string
}

func GetCSTopLayer(msg interface{}) []byte {
	msgname := ""
	if imsg, ok := msg.(IMsg); ok {
		msgname = imsg.GetMsgName()
	}
	b, _ := json.Marshal(msg)
	top := &CS_TopLayer{
		MsgName: msgname,
		Data:    b,
	}
	if top.MsgName[0] == '*' {
		top.MsgName = top.MsgName[1:]
	}
	btop, _ := json.Marshal(top)
	return btop
}

func GetSCTopLayer(msg interface{}) []byte {
	msgname := ""
	if imsg, ok := msg.(IMsg); ok {
		msgname = imsg.GetMsgName()
	}
	b, _ := json.Marshal(msg)
	top := &SC_TopLayer{
		MsgName: msgname,
		Data:    b,
	}
	if top.MsgName[0] == '*' {
		top.MsgName = top.MsgName[1:]
	}
	btop, _ := json.Marshal(top)
	return btop
}

type SC_TopLayer struct {
	MsgName string
	Data    []byte
}

type CS_TopLayer struct {
	MsgName string
	Data    []byte
}

// 账号信息
type AccountInfo struct {
	UUID      string
	LoginName string
}

// 请求登陆账号
type CS_AccountLogin struct {
	LoginName   string
	PassWordMD5 string
}

// 请求登陆账号 的回复
type SC_ResAccountLogin struct {
	// 0 为成功
	Code      int32
	Message   string
	ConnectID string
	Account   *AccountInfo
}

// 请求注册账号
type CS_AccountRegister struct {
	LoginName   string
	PassWordMD5 string
}

// 请求注册账号 的回复
type SC_ResAccountRigster struct {
	// 0 为成功
	Code      int32
	Message   string
	ConnectID string
	Account   *AccountInfo
}

type CS_EnterGame struct {
}

type SC_ResEnterGame struct {
}
