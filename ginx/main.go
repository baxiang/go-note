package main

import (
	"encoding/json"
	"github.com/baxiang/go-note/ginx/mygin"
	"log"
	"net/http"
)

func main() {
	r := mygin.NewEngine()
	r.GET("/ping", func(ctx *mygin.Context) {
		ctx.ResponseWrite.Header().Set("content-type","text/json")
		ctx.ResponseWrite.WriteHeader(http.StatusOK)
		r := map[string]interface{}{"code":0,"message":"ok","data":"pong"}
		data,_:= json.Marshal(r)
		ctx.ResponseWrite.Write(data)
	})
	log.Fatal(r.Run(":8888"))
}
