package client

import "github.com/darkknightjojo/Mytest_Go/chatsever/protocol"

type client interface {
	Dial(address string) error
	Start()
	Close()
	Send(command interface{}) error
	SetName(name string) error
	SendMess(message string) error
	InComing() chan protocol.MessCmd
}
