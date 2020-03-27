package main

import (
	"context"
	hello "github.com/baxiang/go-note/micro-api/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"log"
)

type Say struct {}

var (
	cl hello.SayService
)
func (h *Say)Anything(c *gin.Context){
    c.JSON(200,gin.H{"message":"Hi,this is the Greeter API"})
}
func (h *Say) Hello(c *gin.Context) {
	name := c.Param("name")
	response, err := cl.Hello(context.Background(), &hello.Request{
		Name: name,
	})
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, response)
}


func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
		web.Address(":8080"),
	)

	service.Init()

	cl = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	// setup Greeter Server Client
	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	router.GET("/greeter", say.Anything)
	router.GET("/greeter/:name", say.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
