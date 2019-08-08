package playermodule

import (
	"github.com/liasece/micserver/servercomm"
)

type HandlerClient struct {
	*PlayerModule
}

func (this *HandlerClient) Init(mod *PlayerModule) {
	this.PlayerModule = mod
}

func (this *HandlerClient) OnRecvClientMsg(smsg *servercomm.SForwardFromGate) {
	this.Info("[HandlerClient.OnRecvServerMsg] 收到 Client 消息 %s:%s",
		smsg.MsgName, smsg.GetJson())
}
