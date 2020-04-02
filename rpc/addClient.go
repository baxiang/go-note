package main

import (
	"github.com/baxiang/go-note/rpc/model"
	"log"
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	var r int
	err = client.Call("Calculate.Square", 2, &r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	err = client.Call("Calculate.Add", &model.AddPara {Arga:2,Argb:3}, &r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
