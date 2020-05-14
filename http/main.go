package main

import (
	"fmt"
	"github.com/baxiang/go-note/http/gun"
	"net/http"
)

func IndexHandle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"path = %s\n",r.URL.Path)
}
func HelloHandle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"%v\n",r.Header)

}

func main() {
	r := gun.New()
	r.Get("/",IndexHandle)
	r.Get("/hello",HelloHandle)
	r.Run(":8090")
}
