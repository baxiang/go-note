package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {}

func (engine *Engine)ServeHTTP(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w,"path = %s\n",req.URL.Path)
	case "/hello":
		fmt.Fprintf(w,"%v\n",req.Header)
	default:
		fmt.Fprintf(w,"404 NOT FOUND %s\n",req.URL.Path)
	}
}

func main()  {
	log.Fatal(http.ListenAndServe(":8090",&Engine{}))
}



