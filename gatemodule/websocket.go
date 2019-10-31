package gatemodule

import (
	"encoding/hex"
	"fmt"
	"io"

	"github.com/gobwas/ws"
)

type wsState struct {
	header       ws.Header
	remainUnRead int64
}

type WebSocket struct {
	mod *GatewayModule
}

func (this *WebSocket) Init(mod *GatewayModule) {
	this.mod = mod
}

func (this *wsState) Init() {
}

func (this *WebSocket) Upgrade(conn io.ReadWriter) (ws.Handshake, error) {
	return ws.Upgrade(conn)
}

func (this *WebSocket) DoWrite(writer io.ReadWriter, istate interface{},
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
		this.mod.Error("WebSocket.DoWrite istate.(*wsState) !ok [%+v]",
			istate)
		return 0, istate,
			fmt.Errorf("WebSocket.DoWrite istate.(*wsState) !ok [%+v]",
				istate)
	}
}

func (this *WebSocket) DoRead(reader io.ReadWriter, istate interface{},
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
		this.mod.Info("WebSocket Recv[%s]", hex.EncodeToString(data[:n]))
		return
	} else {
		this.mod.Error("WebSocket.DoRead istate.(*wsState) !ok [%+v]",
			istate)
		return 0, istate,
			fmt.Errorf("WebSocket.DoRead istate.(*wsState) !ok [%+v]",
				istate)
	}
}
