package command

import (
	"github.com/liasece/micserver/msg"
)

var mapping map[string]string

func init() {
	mapping = make(map[string]string)
	ToLogin(&CS_Login{})
	ToLogin(&CS_Register{})
}

func GetServerTypeByID(id uint16) string {
	return GetServerTypeByMsgName(MsgIdToString(id))
}

func GetServerTypeByMsgName(msgname string) string {
	return mapping[msgname]
}

func ToPlayer(m msg.MsgStruct) {
	mapping[m.GetMsgName()] = "player"
}

func ToLogin(m msg.MsgStruct) {
	mapping[m.GetMsgName()] = "login"
}
