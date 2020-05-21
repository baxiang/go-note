package main

import (
	"github.com/baxiang/go-note/http/gun"
	"net/http"
)

func main() {
	//r := gun.New()
	//r.Get("/", func(c *gun.Context) {
	//	c.String(http.StatusOK, "index = %s\n", c.Path)
	//})
	//r.Get("/hello", func(c *gun.Context) {
	//	c.JSON(http.StatusOK,gun.H{"hello":"world"})
	//})
	//r.Run(":8090")

	r :=newRouter()
	r.addRoute("GET","/",nil)
	r.addRoute("GET","/hello/:name",nil)
	r.addRoute("GET","/hi/:name",nil)
	r.addRoute("GET","/hello/a/b",nil)
	r.addRoute("GET","/assets/*filepath",nil)

}
