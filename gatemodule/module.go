package gatemodule

import (
	"ccmd"
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"net"
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
	this.RegAcceptConnect(this.onAcceptConnect)
	this.RegRecvMsg(this.onRecvMsg)
}

func (this *GatewayModule) onAcceptConnect(conn net.Conn) {
	_, err := ws.Upgrade(conn)
	if err != nil {
		this.Error("ws.Upgrade Err[%s]", err.Error())
	} else {
		this.Info("ws.Upgrade")
	}
}

func (this *GatewayModule) onRecvMsg(
	conn *connect.Client, msgbin *msg.MessageBinary) {
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(msgbin.ProtoData, top)
	this.Debug("收到TCP消息 MsgName[%s]", top.MsgName)
	msgname := top.MsgName
	servertype := ccmd.GetServerTypeByMsgName(msgname)
	serverid := conn.GetBindServer(servertype)
	if serverid == "" {
		// 获取一个负载均衡的服务器ID
		serverid = this.GetBalanceServerID(servertype)
		if serverid != "" {
			conn.SetBindServer(servertype, serverid)
		}
	}
	if serverid != "" {
		this.ForwardClientMsgToServer(conn, serverid, 0, msgbin.ProtoData)
	} else {
		this.Error("找不到合适的目标服务器 MsgName[%s] ServerType[%s]",
			msgname, servertype)
	}
}
