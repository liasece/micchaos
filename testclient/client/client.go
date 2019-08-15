package client

import (
	"command"
	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/util"
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

func (this *Client) onConnectRecv(conn *connect.ClientConn,
	msgbinary *msg.MessageBinary) {
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
	conn, err := connect.ClientDial(addr, this.onConnectRecv, nil)
	if err != nil {
		this.Error("connect.ClientDial(%s) err:%s", addr, err.Error())
		return err
	} else {
		this.Debug("链接成功")
	}
	this.Conn = conn

	return nil
}
