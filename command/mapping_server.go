package command

import (
	"github.com/liasece/micserver/msg"
)

var mapping map[string]string

func init() {
	mapping = make(map[string]string)
	ToPlayer(&CS_Login{})
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
