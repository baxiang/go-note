package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.Default()
	mvc.New(app)
	app.Run(iris.Addr(":8080"))
}
