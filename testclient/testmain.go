package main

import (
	"command"
	"fmt"
	"github.com/liasece/micserver/log"
	"time"
)

func main() {
	threadsum := 1
	stopchan := make(chan struct{}, threadsum)
	for i := 0; i < threadsum; i++ {
		client := &Client{
			LoginName: fmt.Sprintf("Jansen%d", i+1),
		}
		client.Logger = log.GetDefaultLogger().Clone()
		client.Logger.SetLogName("client")
		err := client.Dial(":11002")
		if err != nil {
			continue
		}
		client.Conn.SendCmd(&command.CS_Register{
			LoginName:   client.LoginName,
			PassWowdMD5: "psw123456",
		})
		time.Sleep(time.Millisecond * 4)
		stopchan <- struct{}{}
	}
	for i := 0; i < threadsum; i++ {
		<-stopchan
	}
	time.Sleep(time.Millisecond * 400)
	log.CloseLogger()
}
