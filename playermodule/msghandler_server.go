package playermodule

import (
	"github.com/liasece/micserver/servercomm"
)

type HandlerServer struct {
	*PlayerModule
}

func (this *HandlerServer) Init(mod *PlayerModule) {
	this.PlayerModule = mod
}

func (this *HandlerServer) OnForwardToServer(smsg *servercomm.SForwardToServer) {
	this.Info("[HandlerServer.OnForwardToServer] 收到 Server 消息 %d", smsg.MsgID)
}

func (this *HandlerServer) OnForwardFromGate(smsg *servercomm.SForwardFromGate) {
	this.HandlerClient.OnForwardFromGate(smsg)
}
