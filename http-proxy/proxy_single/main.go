package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	//http://127.0.0.1:8001/v1/xxxx/xxx/
	addr := flag.String("addr", "http://127.0.0.1:8001/v1", "server addresss")
	flag.Parse()
	urlTarget, err := url.Parse(*addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(urlTarget)
	log.Fatal(http.ListenAndServe(":8080",proxy))
}
