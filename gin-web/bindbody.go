package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type LoginReq struct {
	Password string `json:"password"  binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/user", func(c *gin.Context) {
		var req LoginReq
		if err :=c.ShouldBindWith(&req,binding.JSON);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"code":400,"mes":"非法请求","data":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{"code":0,"mes":"success","data":req})
	})
	router.Run(":8090")
}
