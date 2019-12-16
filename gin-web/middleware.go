package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Query("token")
		if len(token)>0 {
			c.Set("TOKEN",token)
			c.Next()
		}else {
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限错误","data": struct {}{}})
		}
	}
}


//func main() {
//	router := gin.Default()
//	router.Use(AuthMiddleWare())
//	router.GET("/bar", func(c *gin.Context) {
//		v ,_ := c.Get("TOKEN")
//		c.JSON(http.StatusOK,gin.H{"code":0,"msg":"success","data":v})
//	})
//	router.GET("/foo", func(c *gin.Context) {
//		v ,_ := c.Get("TOKEN")
//		c.JSON(http.StatusOK,gin.H{"code":0,"msg":"success","data":v})
//	})
//	router.Run(":8090")
//}
func main() {
	router := gin.Default()
	router.GET("/bar", func(c *gin.Context) {
		token := c.Query("token")
		if len(token)>0 {
			c.JSON(http.StatusOK,gin.H{"code":0,"msg":"success","data":token})
		}else {
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限错误","data": struct {}{}})
		}
	})
	router.GET("/foo", func(c *gin.Context) {
		token := c.Query("token")
		if len(token)>0 {
			c.JSON(http.StatusOK,gin.H{"code":0,"msg":"success","data":token})
		}else {
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限错误","data": struct {}{}})
		}
	})
	router.Run(":8090")
}