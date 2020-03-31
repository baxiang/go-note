package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DeanThompson/ginpprof"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{"data":"ping"})
	})
	ginpprof.Wrap(router)
	router.Run(":8080")
}
