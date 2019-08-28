package client

import (
	"ccmd"
	"encoding/json"
	"github.com/liasece/micserver/connect"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/msg"
	"github.com/liasece/micserver/util"
)

type Client struct {
	*log.Logger
	Conn       *connect.Client
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

func (this *Client) onConnectRecv(conn *connect.Client,
	msgbinary *msg.MessageBinary) {
	topmsg := &ccmd.SC_TopLayer{}
	json.Unmarshal(msgbinary.ProtoData, topmsg)
	this.Debug("收到消息 %s", topmsg.MsgName)
	if f, ok := this.CmdHandler.mappingFunc[topmsg.MsgName]; ok {
		f(topmsg.Data)
	} else {
		this.Error("未知的消息 MsgName[%s]", topmsg.MsgName)
	}
}

func (this *Client) GetLoginMsg() *ccmd.CS_AccountLogin {
	res := &ccmd.CS_AccountLogin{}
	res.LoginName = this.LoginName
	res.PassWordMD5 = util.HmacSha256ByString(this.Passwd, this.LoginName)
	return res
}

func (this *Client) GetRegsiterMsg() *ccmd.CS_AccountRegister {
	res := &ccmd.CS_AccountRegister{}
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

func (this *Client) SendMsg(msg interface{}) {
	btop := ccmd.GetCSTopLayer(msg)
	this.Conn.SendBytes(0, btop)
}
