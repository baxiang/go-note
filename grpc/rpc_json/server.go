package main

import (
	"github.com/baxiang/go-note/gprc/model"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Operation struct {
}

func (o *Operation) Add(req model.OperationReq, resp *model.OperationResp) error {
	resp.Result = req.A + req.B
	return nil
}

func (o *Operation) Sub(req model.OperationReq, resp *model.OperationResp) error {
	resp.Result = req.A - req.B
	return nil
}

func main() {
	rpc.Register(&Operation{})
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}

		go func(conn net.Conn) { // 并发处理客户端请求
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
