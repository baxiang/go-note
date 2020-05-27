package main

import (
	"context"
	"fmt"
	"github.com/baxiang/go-note/go-rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	client := pb.NewStringServiceClient(conn)
	stringReq := pb.StringRequest{
		A:                    "hello",
		B:                    "world",
	}
	res, err := client.Concat(context.Background(), &stringReq)
	fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, res.Ret)

}
