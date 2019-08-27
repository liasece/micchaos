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

func (this *HandlerServer) OnRecvServerMsg(smsg *servercomm.SForwardToServer) {
	this.Info("[HandlerServer.OnRecvServerMsg] 收到 Server 消息 %d", smsg.MsgID)
}

func (this *HandlerServer) OnRecvGateMsg(smsg *servercomm.SForwardFromGate) {
	this.HandlerClient.OnRecvClientMsg(smsg)
}
