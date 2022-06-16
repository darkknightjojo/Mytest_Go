package main

import (
	"context"
	"flag"
	"fmt"
	valut "github.com/darkknightjojo/Mytest_Go/demo/gokit"
	"github.com/darkknightjojo/Mytest_Go/demo/gokit/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "grpc listen address")
	)

	flag.Parse()
	ctx := context.Background()
	srv := valut.NewService()
	errChan := make(chan error)

	// 捕获终止信号并添加到errchan
	go func() {
		c := make(chan os.Signal, 1)
		// 使用signal.Notify监听SIGINT和SIGTERM信号
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 创建端点实例
	hashEndpoint := valut.MakeHashEndpoint(srv)
	validateEndpoint := valut.MakeValidateEndpoint(srv)

	// 让endpoints成为srv服务器的包装器
	endpoints := valut.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}

	// HTTP transport
	go func() {
		log.Println("http:", *httpAddr)
		handler := valut.NewHttpServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	go func() {
		// 创建TCP侦听器
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler := valut.NewGRPCServer(ctx, endpoints)
		grpcServer := grpc.NewServer()
		pb.RegisterVaultServer(grpcServer, handler)
		errChan <- grpcServer.Serve(listener)
	}()

	log.Fatalln(<-errChan)

}
