package loginmodule

import (
	"command"
	"github.com/liasece/micserver/servercomm"
	"playermodule/boxes"
	"time"
)

type TmpPlayer struct {
	Account *boxes.Account `json:"account"`
}

type HandlerClient struct {
	*LoginModule

	lastCheckTime int64
	msgCount      int64
}

func (this *HandlerClient) Init(mod *LoginModule) {
	this.LoginModule = mod
}

func (this *HandlerClient) OnRecvClientMsg(smsg *servercomm.SForwardFromGate) {
	this.Info("[Login.HandlerClient.OnRecvClientMsg] 收到 Client 消息 %s",
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
	switch smsg.MsgName {
	case "command.CS_Login":
		msg := &command.CS_Login{}
		msg.ReadBinary(smsg.Data)
		this.Info("command.CS_Login: %s", msg.GetJson())
		tmpplayer := &TmpPlayer{}
		err := this.mongo_userinfos.SelectOneByAccount(
			msg.Account, msg.PassWowdMD5, tmpplayer)
		if err != nil {
			// 登陆失败
			this.Error("登陆失败 Err[%s] ReqJson[%s]", err.Error(), msg.GetJson())
		} else {
			// 登陆成功
			this.Info("登陆成功 %s", msg.GetJson())
			send := &command.SC_ResLogin{
				Code:      0,
				Message:   "login secess",
				ConnectID: smsg.ClientConnID,
				Account: &command.AccountInfo{
					UUID: tmpplayer.Account.UUID,
				},
			}
			this.SendServerCmdToServer(smsg.FromServerID, send)
		}
	}
}
