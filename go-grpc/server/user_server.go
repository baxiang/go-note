package main

import (
	"context"
	"github.com/baxiang/go-note/go-grpc/proto"
)

type UserInfoService struct {}

func (user *UserInfoService)GetUserInfo(ctx context.Context,req *proto.UserRequest)

func main() {
	
}
