package boxes

import (
	"ccmd"
)

type Account struct {
	UUID              string `json:"uuid"`
	PhoneNumber       string `json:"phonenumber"`
	PassWordMD5WS     string `json:"passwdmd5"`
	PassWordMD5WSSalt string `json:"passwdmd5salt"`
	LoginName         string `json:"loginname"`
}

func (this *Account) GetMsg() *ccmd.AccountInfo {
	res := &ccmd.AccountInfo{
		LoginName: this.LoginName,
		UUID:      this.UUID,
	}
	return res
}
