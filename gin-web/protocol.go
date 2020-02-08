package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/proto/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello json", "code": 0})
	})

	r.GET("/proto/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hello xml", "code": 0})
	})

	r.GET("/proto/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hello yaml", "code": 0})
	})

	r.GET("/proto/pb", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "ProtoBuf"

		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.GET("/user/info", func(c *gin.Context) {
		user := struct {
			Id int64
			Name string
			IsVIP bool
		}{Id:101,Name:"小明",IsVIP:true}
		response := struct {
			Data interface{}
			Msg string
			Code int
		}{Data:user,Msg:"success",Code:0}

		c.JSON(http.StatusOK, response)
	})

	r.Run(":8090")
}
