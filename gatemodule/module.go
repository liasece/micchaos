package gatemodule

import (
	"command"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/servercomm"
	"github.com/liasece/micserver/tcpconn"
)

type GatewayModule struct {
	module.BaseModule
	testSeqTimes    int64
	testCheckTimeNS int64
	testSwitch      bool
}

func (this *GatewayModule) AfterInitModule() {
	this.BaseModule.AfterInitModule()
	// 当收到客户端发过来的消息时
	gate := this.GetGate()
	if gate != nil {
		gate.RegHandleSocketPackage(this.HandleClientSocketMsg)
	}
	// 当收到服务器间消息时
	subnet := this.GetSubnetManager()
	if subnet != nil {
		subnet.RegHandleServerMsg(this.HandleServerMsg)
	}
}

func (this *GatewayModule) HandleClientSocketMsg(
	conn *tcpconn.ClientConn, msgbin *msg.MessageBinary) {
	this.Debug("收到TCP消息")
	msgname := command.MsgIdToString(msgbin.MessageBinaryHeadL2.CmdID)
	servertype := command.GetServerTypeByMsgName(msgname)
	serverid := conn.Session[servertype]
	if serverid == "" {
		// 获取一个负载均衡的服务器ID
		serverid = this.GetBalanceServerID(servertype)
		if serverid != "" {
			conn.Session[servertype] = serverid
		}
	}
	if serverid != "" {
		this.ForwardClientMsgToServer(conn, serverid, msgname, msgbin.ProtoData)
	}
}

func (this *GatewayModule) HandleServerMsg(smsg *servercomm.SForwardToServer) {
	switch smsg.MsgName {
	case "command.SC_ResLogin":
		{
			msg := &command.SC_ResLogin{}
			msg.ReadBinary(smsg.Data)
			if msg.Code == 0 {
				client := this.GetClientConn(msg.ConnectID)
				if client != nil {
					client.SetVertify(true)
					client.Session["UUID"] = msg.Account.UUID
					this.Info("[gate] 用户登陆成功 %s", msg.GetJson())
				}
			}
			this.SendMsgToClient(this.ModuleID, msg.ConnectID, msg)
		}
	default:
		{
			this.Error("未知消息 %s", smsg.MsgName)
		}
	}
}
