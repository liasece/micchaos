package command

import ()

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
