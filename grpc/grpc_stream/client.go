package main

import (
	"bufio"
	pb "github.com/baxiang/go-note/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"os"
	"log"
	"context"
)

func main() {
	// 创建连接
	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		log.Printf("conn error %v\n", err)
		return
	}
	defer conn.Close()

	// 声明客户端
	client := pb.NewChatClient(conn)
	// 创建双向数据流
	stream, err := client.Hello(context.Background())

	if err != nil {
		log.Printf("创建数据流失败: [%v]\n", err)
	}

	// 启动一个 goroutine 接收命令行输入的指令
	go func() {
		log.Println("请输入消息...")
		buf := bufio.NewReader(os.Stdin)
		for {
			// 获取 命令行输入的字符串， 以回车 \n 作为结束标志
			s, _ := buf.ReadString('\n')
			// 向服务端发送 指令
			if err := stream.Send(&pb.ChatReq{Input: s}); err != nil {
				return
			}
		}
	}()

	for {
		// 接收从 服务端返回的数据流
		rec, err := stream.Recv()
		if err == io.EOF {
			log.Println("收到服务端的结束信号")
			break   //如果收到结束信号，则退出“接收循环”，结束客户端程序
		}

		if err != nil {
			// TODO: 处理接收错误
			log.Println("接收数据出错:", err)
			break
		}
		// 没有错误的情况下，打印来自服务端的消息
		log.Printf("[客户端收到]: %s", rec.Output)
	}
}
