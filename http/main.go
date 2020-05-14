package main

import (
	"github.com/baxiang/go-note/http/gun"
	"net/http"
)

func main() {
	r := gun.New()
	r.Get("/", func(c *gun.Context) {
		c.String(http.StatusOK, "index = %s\n", c.Path)
	})
	r.Get("/hello", func(c *gun.Context) {
		c.JSON(http.StatusOK,gun.H{"hello":"world"})
	})
	r.Run(":8090")
}
