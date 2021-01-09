package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "github.com/baxiang/go-note/grpc_gateway_service/proto"
	"log"
	"net"
)

//定义服务端 实现 约定的接口
type UserInfoService struct{}

//实现 interface
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *proto.UserRequest) (resp *proto.UserResponse, err error) {
	resp = &proto.UserResponse{
		Id:    1568,
		Name:  req.Name,
		Age:   25,
		Title: []string{"Java", "Go"},
	}

	return
}

func main() {
	port := ":6666"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	fmt.Printf("listen %s\n", port)

	s := grpc.NewServer()
	// 将 UserInfoService 注册到 gRPC
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量
	// 需要取地址传参
	proto.RegisterUserInfoServiceServer(s, &UserInfoService{})
	s.Serve(l)
}
