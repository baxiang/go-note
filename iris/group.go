package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	v1 := app.Party("api/v1")
	{
		v1.Post("/login",groupHandle)
	}
	v2 := app.Party("api/v2")
	{
		v2.Post("/login",groupHandle)
	}
	app.Run(iris.Addr(":8080"))

}

func groupHandle(c iris.Context) {
	addr := c.RemoteAddr()
	uri := c.FullRequestURI()
	path := c.Path()
	method := c.Method()
	c.JSON(iris.Map{"addr":addr,"uri":uri,"path":path,"method":method})
}
