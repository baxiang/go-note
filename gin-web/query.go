package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		name:= c.DefaultQuery("name","guest")
		id := c.Query("id")
		c.JSON(http.StatusOK,gin.H{"name":name,"id":id})
	})
    router.Run()
}
