package main

import (
	"command"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/tcpconn"
	"github.com/liasece/micserver/util"
	"io"
	"time"
)

type Client struct {
	*tcpconn.ClientConn
}

func (this *Client) Dial(addr string) error {
	conn, err := tcpconn.ClientDial(addr)
	if err != nil {
		log.Error("tcpconn.ClientDial(%s) err:%s", addr, err.Error())
		return err
	} else {
		log.Info("链接成功")
	}
	this.ClientConn = conn

	netbuffer := util.NewIOBuffer(conn.Conn, 64*1024)
	msgReader := msg.NewMessageBinaryReader(netbuffer)
	go func() {
		for {
			if !this.ClientConn.Check() {
				return
			}
			// 设置阻塞读取过期时间
			err := conn.Conn.SetReadDeadline(
				time.Now().Add(time.Duration(time.Millisecond * 250)))
			if err != nil {
				this.ClientConn.Error("[onNewConnect] SetReadDeadline Err[%s]",
					err.Error())
			}
			// buffer从连接中读取socket数据
			_, err = netbuffer.ReadFromReader()

			// 异常
			if err != nil {
				if err == io.EOF {
					this.ClientConn.Debug("[onNewConnect] "+
						"Scoket数据读写异常,断开连接了,"+
						"scoket返回 Err[%s]", err.Error())
					return
				} else {
					continue
				}
			}

			err = msgReader.RangeMsgBinary(func(msgbinary *msg.MessageBinary) {
				if this.ClientConn.Encryption != msg.EncryptionTypeNone &&
					msgbinary.CmdMask != this.ClientConn.Encryption {
					this.ClientConn.Error("加密方式错误，加密方式应为 %d 此消息为 %d "+
						"MsgID[%d]", this.ClientConn.Encryption,
						msgbinary.CmdMask, msgbinary.CmdID)
				} else {
					// 解析消息
					this.OnRecvSocketPackage(msgbinary)
				}
			})
			if err != nil {
				this.ClientConn.Error("[onNewConnect] 解析消息错误，断开连接 "+
					"Err[%s]", err.Error())
				// 强制移除客户端连接
				return
			}
		}
	}()
	return nil
}

func (this *Client) OnRecvSocketPackage(msgbinary *msg.MessageBinary) {
	log.Info("收到消息 %d", msgbinary.CmdID)
}

// var client *Client

func main() {
	stopchan := make(chan struct{})
	threadsum := 1
	for i := 0; i < threadsum; i++ {
		client := &Client{}
		err := client.Dial(":11002")
		if err != nil {
			continue
		}
		go func() {
			for l := 0; l < threadsum; l++ {
				for i := 0; i < 1; i++ {
					client.SendCmd(&command.CS_Login{
						Account:     "13412341",
						PassWowdMD5: "",
					})
				}
				time.Sleep(time.Millisecond * 4)
			}
			stopchan <- struct{}{}
		}()
	}
	for i := 0; i < threadsum; i++ {
		<-stopchan
	}
	time.Sleep(time.Millisecond * 400)
	log.CloseLogger()
}
