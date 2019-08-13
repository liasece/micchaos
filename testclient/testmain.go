package main

import (
	"fmt"
	"github.com/liasece/micserver/log"
	"testclient/client"
	"time"
)

func main() {
	threadsum := 1
	stopchan := make(chan struct{}, threadsum)
	for i := 0; i < threadsum; i++ {
		c := &client.Client{}
		c.Init(fmt.Sprintf("Jansen%d", i+1), "testpsw99876")
		err := c.Dial(":11002")
		if err != nil {
			continue
		}
		c.Conn.SendCmd(c.GetRegsiterMsg())
		time.Sleep(time.Millisecond * 4)
		stopchan <- struct{}{}
	}
	for i := 0; i < threadsum; i++ {
		<-stopchan
	}
	time.Sleep(time.Millisecond * 400)
	log.CloseLogger()
}
