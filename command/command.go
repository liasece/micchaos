package command

import ()

type CS_AccountLogin struct {
	LoginName   string
	PassWordMD5 string
}

type SC_ResAccountLogin struct {
	// 0 为成功
	Code      int32
	Message   string
	ConnectID string
	Account   *AccountInfo
}

type AccountInfo struct {
	UUID      string
	LoginName string
}

type CS_AccountRegister struct {
	LoginName   string
	PassWordMD5 string
}

type SC_ResAccountRigster struct {
	// 0 为成功
	Code      int32
	Message   string
	ConnectID string
	Account   *AccountInfo
}
