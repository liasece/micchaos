package command

import ()

type CS_Login struct {
	Account     string
	PassWowdMD5 string
}

type SC_ResLogin struct {
	// 0 为成功
	Code      int32
	Message   string
	ConnectID string
	Account   *AccountInfo
}

type AccountInfo struct {
	UUID string
	Name string
}
