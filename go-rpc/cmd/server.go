package main

import (
	"github.com/baxiang/go-note/go-rpc/pb"
	service "github.com/baxiang/go-note/go-rpc/string-service"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterStringServiceServer(server,&service.StringService{})
	server.Serve(listen)

}
