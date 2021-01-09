package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	//多个服务引入多个包
	proto "github.com/baxiang/go-note/grpc_gateway_service/proto"
	"net/http"
)

var userPoint = "localhost:6666"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	//多个服务注册多次即可
	err := proto.RegisterUserInfoServiceHandlerFromEndpoint(ctx, mux, userPoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatal(http.ListenAndServe(":8081", mux))
}
