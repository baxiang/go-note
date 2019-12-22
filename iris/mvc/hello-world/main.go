package main

import (
	"github.com/baxiang/go-note/iris/mvc/hello-world/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)


func main() {
	app := iris.Default()
	mvc.New(app).Handle(&controller.HelloController{})
	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	// http://localhost:8080/custom_path
	app.Run(iris.Addr(":8080"))
}