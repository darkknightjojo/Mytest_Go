package main

import (
	"github.com/darkknightjojo/Mytest_Go/chatsever/server"
	"log"
)

func main() {
	//var s server.Server
	s := server.NewTcpServer()
	err := s.Listen(":3333")
	if err != nil {
		log.Printf("监听端口失败！")
	}
	s.Start()
}
