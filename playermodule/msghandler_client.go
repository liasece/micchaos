package playermodule

import (
	"github.com/liasece/micserver/servercomm"
	"time"
)

type HandlerClient struct {
	*PlayerModule

	lastCheckTime int64
	msgCount      int64
}

func (this *HandlerClient) Init(mod *PlayerModule) {
	this.PlayerModule = mod
}

func (this *HandlerClient) OnRecvClientMsg(smsg *servercomm.SForwardFromGate) {
	this.Info("[HandlerClient.OnRecvClientMsg] 收到 Client 消息 %s",
		smsg.MsgName)
	this.msgCount++
	now := time.Now().UnixNano()
	if now-this.lastCheckTime > 1*1000*1000*1000 {
		this.lastCheckTime = now
		if this.msgCount != 0 {
			this.Error("本秒处理消息 %d", this.msgCount)
		}
		this.msgCount = 0
	}
}
