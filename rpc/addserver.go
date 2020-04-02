package main

import (
	"github.com/baxiang/go-note/rpc/model"
	"net"
	"net/rpc"
	"log"
)

type Calculate struct{}

func (c *Calculate) Square(req int, resp *int) error {
	*resp = req * req
	return nil
}

func (c *Calculate) Add(req model.AddPara, resp *int) error {
	*resp = req.Arga + req.Argb
	return nil
}

func main() {
	rpc.RegisterName("Calculate", new(Calculate))
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
