package main

import (
	"io"
	"net/http"
)

func HelloGo(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"Hello Golang\n")
}

func main() {
	http.HandleFunc("/",HelloGo)
	http.ListenAndServe(":8000",nil)
}
