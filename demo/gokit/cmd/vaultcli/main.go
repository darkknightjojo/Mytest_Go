package main

import (
	"context"
	"flag"
	valut "github.com/darkknightjojo/Mytest_Go/demo/gokit"
	grpcclient "github.com/darkknightjojo/Mytest_Go/demo/gokit/client/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081", "grpc address")
	)
	flag.Parse()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	conn, err := grpc.DialContext(ctx, *grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("grpc dial:", err)
	}

	defer conn.Close()

	vaultService := grpcclient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "hash":
		var password string
		password, args = pop(args)
		hash(ctx, vaultService, password)
	case "validate":
		var password, hash string
		password, args = pop(args)
		hash, args = pop(args)
		validate(ctx, vaultService, password, hash)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func hash(ctx context.Context, service valut.Service, password string) {
	h, err := service.Hash(ctx, password)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(h)
}

func validate(ctx context.Context, service valut.Service, password, hash string) {
	b, err := service.Validate(ctx, password, hash)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(b)
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
