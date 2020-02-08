package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK,gin.H{"user_id":id})
	})

	r.GET("/user/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.JSON(http.StatusOK,gin.H{"id":id,"name":name})
	})

	r.GET("/user/list", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"hello":"world"})
	})
	r.Run()
}
