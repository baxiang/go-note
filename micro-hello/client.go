package main

import (
	"context"
	"fmt"
	proto "github.com/baxiang/go-note/micro-hello/proto"
	"github.com/micro/go-micro"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	g := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := g.Hello(context.Background(),&proto.Request{Name:"micro"} )
	if err != nil {
		fmt.Println(err)
	}
	// Print response
	fmt.Println(rsp.Greeting)
}
