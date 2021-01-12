package main

import (
	"github.com/baxiang/go-note/go-grpc/services"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	services.RegisterProductSvcServer(s, &services.ProductSvc{})
	l, _ := net.Listen("tcp", "8080")
	s.Serve(l)
}
