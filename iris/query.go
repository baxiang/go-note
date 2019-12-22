package main

import "github.com/kataras/iris/v12"

func main() {
	app:= iris.Default()
	app.Get("/hello", func(c iris.Context) {
		firstName:= c.URLParamDefault("firstname", "guest")
		lastName := c.URLParam("lastname")
		c.Writef("Hello %s %s\n",firstName,lastName)
	})
	app.Run(iris.Addr(":8080"))
}
