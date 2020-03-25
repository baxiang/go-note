package main

import (
	"context"
	"github.com/baxiang/go-note/grpc-hello/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct {
}

func (h *HelloService) Hello(c context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Response: "Hello " + r.GetRequest()}, nil
}
func main() {
	server := grpc.NewServer()
	proto.RegisterHelloServiceServer(server, &HelloService{})

	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}
	err = server.Serve(listen)
	if err != nil {
		log.Fatalf(" serve  err: %v", err)
	}
}
