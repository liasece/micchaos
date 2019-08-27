package gatemodule

import (
	"command"
	"encoding/json"
	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/servercomm"
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
	// 当收到服务器间消息时
	subnet := this.GetSubnetManager()
	if subnet != nil {
		subnet.RegHandleServerMsg(this.HandleServerMsg)
	}
}

func (this *GatewayModule) HandleClientSocketMsg(
	conn *connect.ClientConn, msgbin *msg.MessageBinary) {
	this.Debug("收到TCP消息")
	top := &command.CS_TopLayer{}
	json.Unmarshal(msgbin.ProtoData, top)
	msgname := top.MsgName
	servertype := command.GetServerTypeByMsgName(msgname)
	serverid := conn.Session[servertype]
	if serverid == "" {
		// 获取一个负载均衡的服务器ID
		serverid = this.GetBalanceServerID(servertype)
		if serverid != "" {
			conn.Session[servertype] = serverid
		} else {
			this.Error("找不到合适的目标服务器 MsgName[%s] ServerType[%s]",
				msgname, servertype)
		}
	}
	if serverid != "" {
		this.ForwardClientMsgToServer(conn, serverid, 0, msgbin.ProtoData)
	}
}

func (this *GatewayModule) HandleOnNewClient(conn *connect.ClientConn) {
	servertype := util.GetServerIDType(this.ModuleID)
	conn.Session[servertype] = this.ModuleID
	conn.Session["connectid"] = conn.Tempid
}

func (this *GatewayModule) HandleServerMsg(smsg *servercomm.SForwardToServer) {
	switch smsg.MsgID {
	// case command.SC_ResAccountLoginID:
	// 	{
	// 		msg := &command.SC_ResAccountLogin{}
	// 		msg.ReadBinary(smsg.Data)
	// 		if msg.Code == 0 {
	// 			client := this.GetClientConn(msg.ConnectID)
	// 			if client != nil {
	// 				client.SetVertify(true)
	// 				client.Session["UUID"] = msg.Account.UUID
	// 				this.Info("[gate] 用户登陆成功 %s", msg.GetJson())
	// 			}
	// 		}
	// 		this.SendMsgToClient(this.ModuleID, msg.ConnectID, msg)
	// 	}
	default:
		{
			this.Error("未知消息 %d", smsg.MsgID)
		}
	}
}

func (this *GatewayModule) SendMsgToClient(gateid string,
	to string, msg interface{}) {
	btop := command.GetSCTopLayer(msg)
	this.BaseModule.SendBytesToClient(gateid, to, 0, btop)
}
