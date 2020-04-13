package main

import (
	"github.com/baxiang/go-note/gprc/rpc/model"
	"log"
	"net"
	"net/rpc"
	"net/http"
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
	rpc.HandleHTTP()
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	http.Serve(lis, nil)
}
