package main

import (
	pb "github.com/baxiang/go-note/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

//GRPC流 服务端实现流服务


type ChatService struct {
}

func (p *ChatService) Hello(stream  pb.Chat_HelloServer)  error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			// 接收从客户端发来的消息
			rec, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}

			// 如果接收正常，则根据接收到的 字符串 执行相应的指令
			switch rec.Input {
			case "finish\n":
				log.Println("收到'结束对话'指令")
				if err := stream.Send(&pb.ChatRes{Output: "收到结束指令"}); err != nil {
					return err
				}
				// 收到结束指令时，通过 return nil 终止双向数据流
				return nil

			case "back\n":
				log.Println("收到'返回数据流'指令")
				// 收到 收到'返回数据流'指令， 连续返回 10 条数据
				for i := 0; i < 5; i++ {
					if err := stream.Send(&pb.ChatRes{Output: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}

			default:
				// 缺省情况下， 返回 '服务端返回: ' + 输入信息
				log.Printf("[收到客户端消息]: %s", rec.Input)
				if err := stream.Send(&pb.ChatRes{Output: "服务端返回: " + rec.Input}); err != nil {
					return err
				}
			}
		}
	}
}



//注册rpc服务
func main() {
	server := grpc.NewServer()
	pb.RegisterChatServer(server, &ChatService{})
	address, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Printf("Listen err：%s \n", err)
		return
	}
	server.Serve(address)
}
