package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		page := c.Query("page")
		size := c.DefaultQuery("size", "10")
		id := c.PostForm("id")
		name := c.PostForm("name")
		c.JSON(http.StatusOK,gin.H{"Id":id,"page":page,"size":size,"name":name})
	})
	router.Run(":8090")
}