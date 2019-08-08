package boxes

import ()

type Account struct {
	UUID        string `json:"uuid"`
	PhoneNumber string `json:"phonenumber"`
	PassWordMD5 string `json:"passwdmd5"`
}
