package main

import (
	"fmt"
	"github.com/baxiang/go-note/gprc/model"
	"log"
	"net/rpc"
)

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("http error: ", err)
	}

	req := model.OperationReq{A: 5, B: 2}
	var res model.OperationResp

	err = conn.Call("Operation.Add", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("add error: ", err)
	}
	fmt.Printf("%d + %d = %d\n", req.A, req.B, res.Result)

	err = conn.Call("Operation.Sub", req, &res)
	if err != nil {
		log.Fatalln("sub error: ", err)
	}
	fmt.Printf("%d - %d = %d\n", req.A, req.B, res.Result)
}
