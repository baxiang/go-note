package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	var r string
	err = client.Call("HelloService.Hello", "world", &r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
