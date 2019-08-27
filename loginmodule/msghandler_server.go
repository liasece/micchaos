package loginmodule

import (
	"github.com/liasece/micserver/servercomm"
)

type HandlerServer struct {
	*LoginModule
}

func (this *HandlerServer) Init(mod *LoginModule) {
	this.LoginModule = mod
}

func (this *HandlerServer) OnForwardToServer(smsg *servercomm.SForwardToServer) {
	this.Info("[HandlerServer.OnForwardToServer] 收到 Server 消息 %d", smsg.MsgID)
}

func (this *HandlerServer) OnForwardFromGate(smsg *servercomm.SForwardFromGate) {
	this.HandlerClient.OnForwardFromGate(smsg)
}
