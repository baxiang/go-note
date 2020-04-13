package main

import (
	pb "github.com/baxiang/go-note/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"context"
	"fmt"
)

func main() {
	//1.创建与gRPC服务端的连接 grpc.WithInsecure() 建立一个安全连接（跳过了对服务器证书的验证）
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("conn err %s\n", err)
		return
	}
	//2.实例化gRPC客户端
	client := pb.NewOperaServiceClient(conn)
	//3.组装参数
	req := pb.OperaRequest{A: 5, B: 2}
	//4.调用接口
	resp, err := client.Add(context.Background(), &req)
	if err != nil {
		log.Fatalf("add error: %s\n", err)
		return
	}
	fmt.Printf("%d + %d = %d\n", req.GetA(), req.GetB(), resp.GetResult())

	resp, err = client.Sub(context.Background(), &req)
	if err != nil {
		log.Fatalf("sub error %s\n", err)
	}
	fmt.Printf("%d - %d = %d\n", req.A, req.B, resp.GetResult())
}
