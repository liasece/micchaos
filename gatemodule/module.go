package gatemodule

import (
	"ccmd"
	"encoding/json"
	"net"
	"strings"
	"time"

	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/roc"
	"github.com/liasece/micserver/util/monitor"
)

type GatewayModule struct {
	module.BaseModule
	testSeqTimes    int64
	testCheckTimeNS int64
	testSwitch      bool

	// 模块的负载
	ClientMsgLoad          monitor.Load
	lastCheckClientMsgLoad int64

	// websocket 协议
	ws WebSocket
}

func NewGatewayModule(moduleid string) *GatewayModule {
	res := &GatewayModule{}
	res.BaseModule.SetModuleID(moduleid)
	return res
}

// 在 Module 初始化完成之后，注意，此时不一定会连上子网中的其他服务器
func (this *GatewayModule) AfterInitModule() {
	// 调用父类方法
	this.BaseModule.AfterInitModule()

	// 初始化本地
	this.ws.Init(this)
	// 当收到客户端发过来的消息时
	this.HookGate(this)

	// 负载log
	this.TimerManager.RegTimer(time.Second*1, 0, false,
		this.watchClientMsgLoadToLog)
}

// 当收到消息时调用
func (this *GatewayModule) OnRecvClientMsg(
	conn *connect.Client, msgbin *msg.MessageBinary) {
	// 所有客户端的消息都由 CS_TopLayer 包裹
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(msgbin.ProtoData, top)
	this.Debug("收到TCP消息 MsgName[%s]", top.MsgName)

	this.ClientMsgLoad.AddLoad(1)

	msgname := top.MsgName
	servertype := ccmd.GetServerTypeByMsgName(msgname)
	if servertype == "" {
		this.Error("未知消息类型-服务器类型映射 MsgName[%s] ServerType[%s]",
			msgname, servertype)
		return
	}
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

// 在创建新的连接时，将目标连接提升为 websocket 连接
func (this *GatewayModule) OnNewClient(client *connect.Client) {
	if strings.Index(client.RemoteAddr(), "127.0.0.1") < 0 {
		this.Info("尝试升级 websocket 连接 %s", client.RemoteAddr())
		_, err := this.ws.Upgrade(client.IConnection)
		if err != nil {
			this.Error("ws.Upgrade Err[%s]", err.Error())
		} else {
			this.Info("ws.Upgrade")
		}
		// websocket 需要劫持底层的发送及接收流程
		client.HookProtocal(&this.ws)
	}
}

// 关闭客户端连接时触发
func (this *GatewayModule) OnCloseClient(client *connect.Client) {
	this.ROCCallNR(roc.O(ccmd.ROCTypePlayer, client.Session.GetUUID()).F("GateClose"),
		nil)
}

func (this *GatewayModule) OnAcceptClientConnect(conn net.Conn) {
}

func (this *GatewayModule) watchClientMsgLoadToLog(dt time.Duration) bool {
	load := this.ClientMsgLoad.GetLoad()
	incValue := load - this.lastCheckClientMsgLoad
	if incValue > 0 {
		this.Info("[GatewayModule.watchClientMsgLoadToLog] Within %d sec load:[%d]",
			int64(dt.Seconds()), incValue)
	}
	this.lastCheckClientMsgLoad = load
	return true
}
