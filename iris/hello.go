package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	app.Get("/ping", func(c iris.Context) {
		c.JSON(iris.Map{"message":"pong"})
	})
	app.Run(iris.Addr(":8080"))
}
