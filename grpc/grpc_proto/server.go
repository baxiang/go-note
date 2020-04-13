package main

import (
	"context"
	"fmt"
	pb "github.com/baxiang/go-note/grpc/proto"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	"net"
)

type OperaService struct {
}

// 业务逻辑代码
func (a *OperaService) Add(ctx context.Context, req *pb.OperaRequest) (res *pb.OperaResponse, err error) {
	res = &pb.OperaResponse{
		Result: req.A + req.B,
	}
	return
}

func (a *OperaService) Sub(ctx context.Context, req *pb.OperaRequest) (res *pb.OperaResponse, err error) {
	res = &pb.OperaResponse{
		Result: req.A - req.B,
	}
	return
}
func main() {
	//1.监听
	list, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Printf("list err：%v \n", err)
		return
	}
	//2、实例化gRPC
	s := grpc.NewServer()
	//3.在gRPC上注册微服务
	//第二个参数接口类型的变量"errors"
	pb.RegisterOperaServiceServer(s, &OperaService{})

	//reflection.Register(s)////关闭注册反射
	//4.启动gRPC服务端
	s.Serve(list)
}
