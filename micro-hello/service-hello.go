package main

import (
	"context"
	"github.com/baxiang/go-note/micro-hello/hello"
	"github.com/micro/go-micro"
	"log"
)

type Say struct {}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("go.micro.srv.greeter"),
	)
	service.Init()
	hello.RegisterSayHandler(service.Server(),new(Say))
	if err :=service.Run();err!=nil {
		log.Fatal()
	}
}
