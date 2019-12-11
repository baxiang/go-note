package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello",
		func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"hello world")
	})
	log.Fatal(http.ListenAndServe(":8080",nil))
}
