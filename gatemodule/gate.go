package gatemodule

import (
	"command"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/servercomm"
	"github.com/liasece/micserver/tcpconn"
	"time"
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
	// 业务逻辑
	if this.GetModuleID() == "gate001" {
		time.AfterFunc(2*time.Second, func() {
			for i := 0; i < 2000; i++ {
				// t := command.TetstValue
				// this.BroadcastServerCmd(&t)
				send := &command.Test{Seq: msg.MessageMaxSize - 100}
				send.Data = make([]byte, send.Seq)
				for i, _ := range send.Data {
					send.Data[i] = byte(i)
				}
				this.BroadcastServerCmd(send)
			}
		})
	}
}

func (this *GatewayModule) HandleClientSocketMsg(
	conn *tcpconn.ClientConn, msgbin *msg.MessageBinary) {
	this.Debug("收到TCP消息")
}

func (this *GatewayModule) HandleServerMsg(
	conn *tcpconn.ServerConn, smsg *servercomm.SForwardToServer) {
	// this.Debug("收到服务器消息 %s", smsg.MsgName)
	if smsg.MsgName == "command.Test" {
		msg := &command.Test{}
		msg.ReadBinary(smsg.Data)

		// srcJson := command.TetstValue.GetJson()
		// dirJson := msg.GetJson()
		// if srcJson != dirJson {
		// 	this.Error("json比较不正确：\n%s\n%s\n%x", srcJson, dirJson, smsg.Data)
		// }
		// seq := 1
		if msg.Seq != int64(len(msg.Data)) {
			this.Error("错误的长度 Seq:%d LenData:%d", msg.Seq, len(msg.Data))
			return
		}
		for i, v := range msg.Data {
			if v != byte(i) {
				this.Error("错误的数据 i:%d v:%d bytei:%d", i, v, byte(i))
				return
			}
		}
		seq := msg.Seq
		this.testSeqTimes++
		nowNs := time.Now().UnixNano()
		if nowNs-this.testCheckTimeNS > 1000*1000*1000 {
			this.testCheckTimeNS = nowNs
			this.Info("%s seq:%d 测试接受包数量: %d", this.GetModuleID(),
				seq, this.testSeqTimes)
			this.testSeqTimes = 0
		}
		// t := command.TetstValue
		// this.BroadcastServerCmd(&t)
		send := &command.Test{Seq: (seq)}
		send.Data = make([]byte, send.Seq)
		for i, _ := range send.Data {
			send.Data[i] = byte(i)
		}
		this.BroadcastServerCmd(send)
	} else {
		this.Error("未知消息 %s", smsg.MsgName)
	}
}
