package loginmodule

import (
	"github.com/liasece/micserver/servercomm"
	"github.com/liasece/micserver/session"
)

type HandlerServer struct {
	*LoginModule
}

func (this *HandlerServer) Init(mod *LoginModule) {
	this.LoginModule = mod
}

func (this *HandlerServer) OnModuleMessage(smsg *servercomm.ModuleMessage) {
	this.Info("[HandlerServer.OnModuleMessage] 收到 Module 消息 %d", smsg.MsgID)
}

func (this *HandlerServer) OnClientMessage(session *session.Session,
	smsg *servercomm.ClientMessage) {
	this.HandlerClient.OnClientMessage(session, smsg)
}
