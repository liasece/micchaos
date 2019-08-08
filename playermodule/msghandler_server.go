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

func (this *HandlerServer) OnRecvServerMsg(smsg *servercomm.SForwardToServer) {
	this.Info("[HandlerServer.OnRecvServerMsg] 收到 Server 消息 %s", smsg.MsgName)
}

func (this *HandlerServer) OnRecvGateMsg(smsg *servercomm.SForwardFromGate) {
	this.HandlerClient.OnRecvClientMsg(smsg)
}
