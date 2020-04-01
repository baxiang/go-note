package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello http\n"))
	})
	log.Fatal(http.ListenAndServe(":8090",nil))
}
