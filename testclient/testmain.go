package main

import (
	"fmt"
	"github.com/liasece/micserver/log"
	"testclient/client"
	"time"
)

func run(ch chan struct{}, i int) {
	defer func() {
		ch <- struct{}{}
	}()

	c := &client.Client{}
	c.Init(fmt.Sprintf("Jansen%d", i+1), "testpsw99876")
	err := c.Dial(":11002")
	if err != nil {
		return
	}
	for i := 0; i < 100000; i++ {
		c.SendMsg(c.GetRegsiterMsg())
	}
	time.Sleep(time.Second * 30)
}

func main() {
	threadsum := 10
	stopchan := make(chan struct{}, threadsum)
	for i := 0; i < threadsum; i++ {
		go run(stopchan, i)
	}
	for i := 0; i < threadsum; i++ {
		<-stopchan
	}
	time.Sleep(time.Millisecond * 400)
	log.CloseLogger()
}
