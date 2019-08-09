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
	*log.Logger
	Conn      *tcpconn.ClientConn
	LoginName string
}

func (this *Client) OnRecvSocketPackage(msgbinary *msg.MessageBinary) {
	this.Debug("收到消息 %d", msgbinary.CmdID)
	switch msgbinary.CmdID {
	case command.SC_ResLoginID:
		msg := &command.SC_ResLogin{}
		msg.ReadBinary(msgbinary.ProtoData)
		this.OnResLogin(msg)
	case command.SC_ResRigsterID:
		msg := &command.SC_ResRigster{}
		msg.ReadBinary(msgbinary.ProtoData)
		this.OnResRigster(msg)
	}
}

func (this *Client) OnResRigster(msg *command.SC_ResRigster) {
	this.Conn.SendCmd(&command.CS_Login{
		LoginName:   this.LoginName,
		PassWowdMD5: "psw123456",
	})
	if msg.Code != 0 {
		this.Error("注册账号失败 %s", msg.GetJson())
		return
	}
	this.Info("注册成功 %s", msg.GetJson())
	if msg.Account.LoginName != "" {
		this.Logger.SetLogName(msg.Account.LoginName)
	}
}

func (this *Client) OnResLogin(msg *command.SC_ResLogin) {
	if msg.Code != 0 {
		this.Error("登陆失败 %s", msg.GetJson())
		return
	}
	this.Info("登陆成功 %s", msg.GetJson())
	if msg.Account.LoginName != "" {
		this.Logger.SetLogName(msg.Account.LoginName)
	}
}

func (this *Client) Dial(addr string) error {
	conn, err := tcpconn.ClientDial(addr)
	if err != nil {
		this.Error("tcpconn.ClientDial(%s) err:%s", addr, err.Error())
		return err
	} else {
		this.Debug("链接成功")
	}
	this.Conn = conn

	go this.recvMsgProcess()
	return nil
}

func (this *Client) recvMsgProcess() {
	netbuffer := util.NewIOBuffer(this.Conn.Conn, 64*1024)
	msgReader := msg.NewMessageBinaryReader(netbuffer)
	for {
		if !this.Conn.Check() {
			return
		}
		// 设置阻塞读取过期时间
		err := this.Conn.Conn.SetReadDeadline(
			time.Now().Add(time.Duration(time.Millisecond * 250)))
		if err != nil {
			this.Error("[recvMsgProcess] SetReadDeadline Err[%s]",
				err.Error())
		}
		// buffer从连接中读取socket数据
		_, err = netbuffer.ReadFromReader()

		// 异常
		if err != nil {
			if err == io.EOF {
				this.Debug("[recvMsgProcess] "+
					"Scoket数据读写异常,断开连接了,"+
					"scoket返回 Err[%s]", err.Error())
				return
			} else {
				continue
			}
		}

		err = msgReader.RangeMsgBinary(func(msgbinary *msg.MessageBinary) {
			// 解析消息
			this.OnRecvSocketPackage(msgbinary)
		})
		if err != nil {
			this.Error("[recvMsgProcess] 解析消息错误，断开连接 "+
				"Err[%s]", err.Error())
			// 强制移除客户端连接
			return
		}
	}
}
