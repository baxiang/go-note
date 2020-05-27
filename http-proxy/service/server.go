package  main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type serverHandle struct {}

func (s *serverHandle) indexHandler(w http.ResponseWriter, req *http.Request) {
	path := fmt.Sprintf("path =%s\n", req.URL.Path)
	clientIP := fmt.Sprintf("RemoteAddr=%s\n", req.RemoteAddr)
	header:=fmt.Sprintf("query =%v\n",req.URL.Query())
	io.WriteString(w, path)
	io.WriteString(w, clientIP)
	io.WriteString(w, header)
}
func (s *serverHandle) ErrorHandler(w http.ResponseWriter, req *http.Request) {
	upath := "error handler\n"
	w.WriteHeader(500)
	io.WriteString(w, upath)
}

func (s *serverHandle) TimeoutHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(4*time.Second)
	upath := "timeout handler\n"
	w.WriteHeader(200)
	io.WriteString(w, upath)
}
func main(){
	port := flag.String("port","8001","server port")
	flag.Parse()
	h := &serverHandle{}
	mux := http.NewServeMux()
	mux.HandleFunc("/",h.indexHandler)
	mux.HandleFunc("/test/error",h.ErrorHandler)
	mux.HandleFunc("/foo/test/timeout",h.TimeoutHandler)
	s :=http.Server{
		Addr:              ":"+*port,
		Handler:           mux,
		WriteTimeout:      time.Second*3,
	}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()
	quit := make(chan os.Signal)
	go func() {
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	}()
	<-quit
}