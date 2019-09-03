package gatemodule

import (
	"ccmd"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"io"
	"time"
)

type GatewayModule struct {
	module.BaseModule
	testSeqTimes    int64
	testCheckTimeNS int64
	testSwitch      bool

	lastCheckTime int64
	msgCount      int64
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
	// 当收到客户端发过来的消息时
	// this.RegOnNewClient(this.onNewClient)
	this.RegOnRecvClientMsg(this.onRecvClientMsg)
}

// 当收到消息时调用
func (this *GatewayModule) onRecvClientMsg(
	conn *connect.Client, msgbin *msg.MessageBinary) {
	// 所有客户端的消息都由 CS_TopLayer 包裹
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(msgbin.ProtoData, top)
	this.Debug("收到TCP消息 MsgName[%s]", top.MsgName)

	this.msgCount++
	now := time.Now().UnixNano()
	if now-this.lastCheckTime > 1*1000*1000*1000 {
		this.lastCheckTime = now
		if this.msgCount != 0 {
			this.Error("本秒处理消息 %d", this.msgCount)
		}
		this.msgCount = 0
	}

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
func (this *GatewayModule) onNewClient(client *connect.Client) {
	_, err := ws.Upgrade(client.Conn)
	if err != nil {
		this.Error("ws.Upgrade Err[%s]", err.Error())
	} else {
		this.Info("ws.Upgrade")
	}
	// websocket 需要劫持底层的发送及接收流程
	client.RegDoReadTCPBytes(this.doReadWSBytes)
	client.RegDoSendTCPBytes(this.doSendWSBytes)
}

type wsState struct {
	header       ws.Header
	remainUnRead int64
}

func (this *wsState) Init() {
}

func (this *GatewayModule) doSendWSBytes(writer io.ReadWriter, istate interface{},
	data []byte) (n int, resstate interface{}, err error) {
	if istate == nil {
		newstate := &wsState{}
		newstate.Init()
		istate = newstate
	}
	if state, ok := istate.(*wsState); ok {
		// Reset the Masked flag, server frames must not be masked as
		// RFC6455 says.
		// 取上次读到的头部信息
		header := state.header
		// 无计算掩码
		header.Masked = false
		// 长度
		header.Length = int64(len(data))
		if err = ws.WriteHeader(writer, header); err != nil {
			// handle error
			return
		}
		if n, err = writer.Write(data); err != nil {
			// handle error
			return
		}
		// 关闭 ws 层连接
		if header.OpCode == ws.OpClose {
			err = io.EOF
		}
		return
	} else {
		this.Error("GatewayModule.doSendWSBytes istate.(*wsState) !ok [%+v]",
			istate)
		return 0, istate,
			fmt.Errorf("GatewayModule.doSendWSBytes istate.(*wsState) !ok [%+v]",
				istate)
	}
}

func (this *GatewayModule) doReadWSBytes(reader io.ReadWriter, istate interface{},
	data []byte) (n int, resstate interface{}, err error) {
	if istate == nil {
		newstate := &wsState{}
		newstate.Init()
		istate = newstate
	}
	resstate = istate
	if state, ok := istate.(*wsState); ok {
		if state.remainUnRead <= 0 {
			// 上一帧没有未读数据，开始读新的一帧
			state.header, err = ws.ReadHeader(reader)
			if err != nil {
				// ReadHeader err
				return
			}
			state.remainUnRead = state.header.Length
		}
		// 如果最大可读大小大于剩余未读数据大小，
		// 需要限制data的长度防止读到下一帧帧头
		readsize := int64(len(data))
		if readsize > state.remainUnRead {
			readsize = state.remainUnRead
		}
		// 读数据
		n, err = reader.Read(data[:readsize])
		// 剩余未读数据需要减少本次读取的大小
		state.remainUnRead -= int64(n)
		if err != nil {
			// io.Read err
			return
		}
		if state.header.Masked {
			ws.Cipher(data, state.header.Mask, 0)
		}
		this.Info("WebSocket Recv[%s]", hex.EncodeToString(data[:n]))
		return
	} else {
		this.Error("GatewayModule.doReadWSBytes istate.(*wsState) !ok [%+v]",
			istate)
		return 0, istate,
			fmt.Errorf("GatewayModule.doReadWSBytes istate.(*wsState) !ok [%+v]",
				istate)
	}
}
