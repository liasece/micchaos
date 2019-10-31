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

func (this *HandlerServer) OnServerMessage(smsg *servercomm.ServerMessage) {
	this.Info("[HandlerServer.OnServerMessage] 收到 Server 消息 %d", smsg.MsgID)
}

func (this *HandlerServer) OnClientMessage(smsg *servercomm.ClientMessage) {
	this.HandlerClient.OnClientMessage(smsg)
}
