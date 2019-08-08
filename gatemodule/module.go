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
	serverid := conn.BindServerID[servertype]
	if serverid == "" {
		// 获取一个负载均衡的服务器ID
		serverid = this.GetBalanceServerID(servertype)
		if serverid != "" {
			conn.BindServerID[servertype] = serverid
		}
	}
	if serverid != "" {
		this.SendGateMsgByTmpID(conn, serverid, msgname, msgbin.ProtoData)
	}
}

func (this *GatewayModule) HandleServerMsg(smsg *servercomm.SForwardToServer) {
	if smsg.MsgName == "command.Test" {
	} else {
		this.Error("未知消息 %s", smsg.MsgName)
	}
}
