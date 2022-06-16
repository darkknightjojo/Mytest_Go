package main

import (
	"context"
	"fmt"
	"github.com/darkknightjojo/Mytest_Go/Test/gRPC/protocol"
	"google.golang.org/grpc"
	"net"
)

const (
	post = "127.0.0.1:18887"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *protocol.HelloRequest) (*protocol.HelloReplay, error) {
	return &protocol.HelloReplay{Message: "hello" + in.GetName()}, nil
}

func (s *server) GetHelloMsg(ctx context.Context, in *protocol.HelloRequest) (*protocol.HelloMessage, error) {
	return &protocol.HelloMessage{
		Msg: "this is from server",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", post)
	if err != nil {
		fmt.Println("网络异常", err)
	}

	ser := grpc.NewServer()
	protocol.RegisterHelloServerServer(ser, &server{})
	err = ser.Serve(listener)

	if err != nil {
		fmt.Println("网络异常", err)
	}
}
