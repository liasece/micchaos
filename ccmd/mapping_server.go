package ccmd

import (
	"reflect"
)

var mapping map[string]string

func init() {
	mapping = make(map[string]string)

	ToLogin(&CS_AccountLogin{})
	ToLogin(&CS_AccountRegister{})

	ToPlayer(&CS_EnterGame{})
}

func GetModuleTypeByMsgName(msgname string) string {
	return mapping[msgname]
}

func ToPlayer(m interface{}) {
	msgname := reflect.TypeOf(m).String()
	if msgname[0] == '*' {
		msgname = msgname[1:]
	}
	mapping[msgname] = "player"
}

func ToLogin(m interface{}) {
	msgname := reflect.TypeOf(m).String()
	if msgname[0] == '*' {
		msgname = msgname[1:]
	}
	mapping[msgname] = "login"
}
