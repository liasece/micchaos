package command

import ()

type CS_Login struct {
	LoginName   string
	PassWordMD5 string
}

type SC_ResLogin struct {
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

type CS_Register struct {
	LoginName   string
	PassWordMD5 string
}

type SC_ResRigster struct {
	// 0 为成功
	Code      int32
	Message   string
	ConnectID string
	Account   *AccountInfo
}
