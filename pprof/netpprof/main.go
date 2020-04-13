package main

import (
	"net/http"
	"log"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
