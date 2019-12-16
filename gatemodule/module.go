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
	"github.com/liasece/micserver/rocutil"
	"github.com/liasece/micserver/rocutil/options"
	"github.com/liasece/micserver/util/monitor"
)

type GatewayModule struct {
	module.BaseModule

	testSeqTimes    int64
	testCheckTimeNS int64
	testSwitch      bool
	// 模块的负载
	clientMsgLoad          monitor.Load
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

	// 初始化本地ROC对象
	_, err := rocutil.ServerROCObj(this, this, roc.ROCObjType("gatemodule"),
		this.GetModuleID(), &options.Options{
			CheckFuncName: func(name string) (string, bool) {
				if name[:4] != "ROC_" {
					return "", false
				}
				name = name[4:]
				this.Info("ROC注册调用函数:%s", name)
				return name, true
			},
		})
	if err != nil {
		this.Error("rocutil.ServerROCObj Err[%s]", err.Error())
	}

	// 负载log
	this.RegTimer(time.Second*1, 0, false,
		this.watchClientMsgLoadToLog)
}

// 当收到消息时调用
func (this *GatewayModule) OnRecvClientMsg(
	conn *connect.Client, msgbin *msg.MessageBinary) {
	// 所有客户端的消息都由 CS_TopLayer 包裹
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(msgbin.ProtoData, top)
	this.Debug("收到TCP消息 MsgName[%s]", top.MsgName)

	this.clientMsgLoad.AddLoad(1)

	msgname := top.MsgName
	moduletype := ccmd.GetModuleTypeByMsgName(msgname)
	if moduletype == "" {
		this.Error("未知消息类型-服务器类型映射 MsgName[%s] ServerType[%s]",
			msgname, moduletype)
		return
	}
	moduleid := conn.GetBind(moduletype)
	if moduleid == "" {
		// 获取一个负载均衡的服务器ID
		moduleid = this.GetBalanceModuleID(moduletype)
		if moduleid != "" {
			conn.SetBind(moduletype, moduleid)
		}
	}
	if moduleid != "" {
		this.ForwardClientMsgToModule(conn, moduleid, 0, msgbin.ProtoData)
	} else {
		this.Error("找不到合适的目标服务器 MsgName[%s] ModuleType[%s]",
			msgname, moduletype)
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
	this.ROCCallNR(
		roc.O(ccmd.ROCTypePlayer, client.Session.GetUUID()).F("GateClose"),
		nil)
}

func (this *GatewayModule) OnAcceptClientConnect(conn net.Conn) {
}

func (this *GatewayModule) watchClientMsgLoadToLog(dt time.Duration) bool {
	load := this.clientMsgLoad.GetLoad()
	incValue := load - this.lastCheckClientMsgLoad
	if incValue > 0 {
		this.Info("[GatewayModule.watchClientMsgLoadToLog] Within %d sec load:[%d]",
			int64(dt.Seconds()), incValue)
	}
	this.lastCheckClientMsgLoad = load
	return true
}

type ROCUtilTest struct {
	Str     string
	Int     int
	Float32 float32
}

func (this *GatewayModule) ROC_ShowInfo(str string, load *ROCUtilTest) {
	this.Info("ShowInfo %+v %+v", str, load)
}
