package main

import (
	"context"
	"github.com/baxiang/go-note/grpc-hello/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := proto.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &proto.HelloRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client err: %v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}
