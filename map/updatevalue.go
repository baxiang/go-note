package main

import "fmt"

func main() {
	m := map[string]int{"one": 1}
	fmt.Println("key one before ",m["one"])
	modifyMap(m)
	fmt.Println("key one after ",m["one"])
}

func modifyMap(m map[string]int) {
	m["one"] = 2
}