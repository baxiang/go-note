package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	app.Post("/form", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamInt32Default("page", 0)
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		ctx.JSON(iris.Map{"id":id,"page":page,"name":name,"message":message})
	})

	app.Run(iris.Addr(":8080"))
}