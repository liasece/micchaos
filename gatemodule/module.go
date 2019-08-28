package gatemodule

import (
	"ccmd"
	"encoding/json"
	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/util"
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
		gate.RegOnNewConn(this.HandleOnNewClient)
	}
}

func (this *GatewayModule) HandleClientSocketMsg(
	conn *connect.ClientConn, msgbin *msg.MessageBinary) {
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(msgbin.ProtoData, top)
	this.Debug("收到TCP消息 MsgName[%s]", top.MsgName)
	msgname := top.MsgName
	servertype := ccmd.GetServerTypeByMsgName(msgname)
	serverid := conn.Session.GetBindServer(servertype)
	if serverid == "" {
		// 获取一个负载均衡的服务器ID
		serverid = this.GetBalanceServerID(servertype)
		if serverid != "" {
			conn.Session.SetBindServer(servertype, serverid)
		}
	}
	if serverid != "" {
		this.ForwardClientMsgToServer(conn, serverid, 0, msgbin.ProtoData)
	} else {
		this.Error("找不到合适的目标服务器 MsgName[%s] ServerType[%s]",
			msgname, servertype)
	}
}

func (this *GatewayModule) HandleOnNewClient(conn *connect.ClientConn) {
	servertype := util.GetServerIDType(this.ModuleID)
	conn.Session.SetBindServer(servertype, this.ModuleID)
}
