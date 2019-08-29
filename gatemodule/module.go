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
	this.RegNewClient(this.onNewClient)
	this.RegRecvMsg(this.onRecvMsg)
}

func (this *GatewayModule) onNewClient(client *connect.Client) {
	_, err := ws.Upgrade(client.Conn)
	if err != nil {
		this.Error("ws.Upgrade Err[%s]", err.Error())
	} else {
		this.Info("ws.Upgrade")
	}
	client.RegReadTCPBytes(this.doReadWSBytes)
	client.RegSendTCPBytes(this.doSendWSBytes)
}

func (this *GatewayModule) onRecvMsg(
	conn *connect.Client, msgbin *msg.MessageBinary) {
	top := &ccmd.CS_TopLayer{}
	json.Unmarshal(msgbin.ProtoData, top)
	this.Debug("收到TCP消息 MsgName[%s]", top.MsgName)
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
