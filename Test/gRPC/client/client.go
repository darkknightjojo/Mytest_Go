package main

import (
	"context"
	"fmt"
	"github.com/darkknightjojo/Mytest_Go/Test/gRPC/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const post = "127.0.0.1:18887"

func main() {
	conn, err := grpc.Dial(post, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接服务器失败", err)
	}

	defer conn.Close()

	client := protocol.NewHelloServerClient(conn)
	r1, err := client.SayHello(context.Background(), &protocol.HelloRequest{Name: "WYD"})
	if err != nil {
		fmt.Println("could not get hello1 from server", err)
		return
	}
	fmt.Println("HelloServer resp: ", r1.Message)

	r2, err := client.GetHelloMsg(context.Background(), &protocol.HelloRequest{Name: "WYD"})
	if err != nil {
		fmt.Println("could not get hello2 from server", err)
		return
	}
	fmt.Println("HelloServer resp: ", r2.Msg)
}
