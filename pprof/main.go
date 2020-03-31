package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	_ "net/http/pprof"
)

type request struct {
	Agent string `json:"agent"`
	Time int64 `json:"time"`
	Method  string `json:"method"`
	Path  string `json:"path"`
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		result := request{
			Agent: r.UserAgent(),
			Time:  time.Now().Unix(),
			Method:   r.Method,
			Path: r.URL.Path,
		}
		data,_:= json.Marshal(&result)
		fmt.Fprint(w,string(data))
	})
	if err := http.ListenAndServe(":8080", nil);err!=nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
