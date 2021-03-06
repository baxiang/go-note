package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {

}

func (h *HelloService)Hello(request string,response *string)error{
	*response ="hello"+request
	return nil
}

func main() {
	rpc.RegisterName("HelloService",new(HelloService))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for{
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		fmt.Println("conn")
		go rpc.ServeConn(conn)
	}
}
