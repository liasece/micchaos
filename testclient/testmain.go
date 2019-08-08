package main

import (
	"command"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/tcpconn"
	"time"
)

type Client struct {
	*tcpconn.ClientConn
}

func (this *Client) Dial(addr string) {
	conn, err := tcpconn.ClientDial(addr)
	if err != nil {
		log.Error("tcpconn.ClientDial(%s) err:%s", addr, err.Error())
		return
	} else {
		log.Info("链接成功")
	}
	this.ClientConn = conn
}

var client *Client

func main() {
	client = &Client{}
	client.Dial(":11002")
	client.SendCmd(&command.CS_Login{
		Account: "18378396103",
	})
	time.Sleep(time.Millisecond * 400)
	log.CloseLogger()
}
