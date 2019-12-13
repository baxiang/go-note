package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":0,
				"message": "success",
				"data":"pong",
			})
		})
		r.Run()

}
