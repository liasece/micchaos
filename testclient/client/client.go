package client

import (
	"command"
	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/util"
	"io"
	"time"
)

type Client struct {
	*log.Logger
	Conn       *connect.ClientConn
	LoginName  string
	Passwd     string
	CmdHandler CmdHandler
}

func (this *Client) Init(name, passwd string) {
	this.LoginName = name
	this.Passwd = passwd
	this.Logger = log.GetDefaultLogger().Clone()
	this.Logger.SetLogName("client")
	this.CmdHandler.Init(this)
}

func (this *Client) OnRecvSocketPackage(msgbinary *msg.MessageBinary) {
	msgname := command.MsgIdToString(msgbinary.CmdID)
	this.Debug("收到消息 %s", msgname)
	if f, ok := this.CmdHandler.mappingFunc[msgname]; ok {
		f(msgbinary)
	} else {
		this.Error("未知的消息 %d:%s", msgbinary.CmdID, msgname)
	}
}

func (this *Client) GetLoginMsg() *command.CS_AccountLogin {
	res := &command.CS_AccountLogin{}
	res.LoginName = this.LoginName
	res.PassWordMD5 = util.HmacSha256ByString(this.Passwd, this.LoginName)
	return res
}

func (this *Client) GetRegsiterMsg() *command.CS_AccountRegister {
	res := &command.CS_AccountRegister{}
	res.LoginName = this.LoginName
	res.PassWordMD5 = util.HmacSha256ByString(this.Passwd, this.LoginName)
	return res
}

func (this *Client) Dial(addr string) error {
	conn, err := connect.ClientDial(addr)
	if err != nil {
		this.Error("connect.ClientDial(%s) err:%s", addr, err.Error())
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
