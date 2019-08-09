package boxes

import (
	"command"
)

type Account struct {
	UUID        string `json:"uuid"`
	PhoneNumber string `json:"phonenumber"`
	PassWordMD5 string `json:"passwdmd5"`
	LoginName   string `json:"loginname"`
}

func (this *Account) GetMsg() *command.AccountInfo {
	res := &command.AccountInfo{
		LoginName: this.LoginName,
		UUID:      this.UUID,
	}
	return res
}
