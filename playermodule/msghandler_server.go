package playermodule

import (
	"github.com/liasece/micserver/servercomm"
	"github.com/liasece/micserver/session"
)

type HandlerServer struct {
	*PlayerModule
}

func (this *HandlerServer) Init(mod *PlayerModule) {
	this.PlayerModule = mod
}

func (this *HandlerServer) OnModuleMessage(smsg *servercomm.ModuleMessage) {
	this.Info("[HandlerServer.OnModuleMessage] 收到 Module 消息 %d", smsg.MsgID)
}

func (this *HandlerServer) OnClientMessage(session *session.Session,
	smsg *servercomm.ClientMessage) {
	this.HandlerClient.OnClientMessage(session, smsg)
}
