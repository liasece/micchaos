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

func (this *Client) Dial(addr string) error {
	conn, err := tcpconn.ClientDial(addr)
	if err != nil {
		log.Error("tcpconn.ClientDial(%s) err:%s", addr, err.Error())
		return err
	} else {
		log.Info("链接成功")
	}
	this.ClientConn = conn
	return nil
}

// var client *Client

func main() {
	stopchan := make(chan struct{})
	threadsum := 100
	for i := 0; i < threadsum; i++ {
		client := &Client{}
		err := client.Dial(":11002")
		if err != nil {
			continue
		}
		go func() {
			for l := 0; l < 500; l++ {
				for i := 0; i < 500; i++ {
					client.SendCmd(&command.CS_Login{
						Account: "18378396103",
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
