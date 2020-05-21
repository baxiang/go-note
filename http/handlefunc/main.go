package main

import (
	"fmt"
	"log"
	"net/http"
)
func IndexHandle(w http.ResponseWriter,r *http.Request){
	  fmt.Fprintf(w,"path = %s\n",r.URL.Path)
}
func HelloHandle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"%v\n",r.Header)

}

func main() {
	http.HandleFunc("/",IndexHandle)
	http.HandleFunc("/hello",HelloHandle)
	log.Fatal(http.ListenAndServe(":8090",nil))
}
