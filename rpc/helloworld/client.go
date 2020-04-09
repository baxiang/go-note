package main

import (
	"fmt"
	"net/rpc"
	"log"
)

func main() {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
