package command

import ()

type CS_Login struct {
	Account string
}

type SC_ResLogin struct {
	Code     int32
	UserInfo *UserInfo
}

type UserInfo struct {
	UUID string
	Name string
}
