package main

import "fmt"

func main() {
	m := map[string]int{"one": 1}
	if v,ok := m["one"];ok{
		fmt.Println("key one  exist",v)
	}
	if v,ok := m["two"];!ok {
		fmt.Println("key two  not exist",v)
	}
}

func modifyMap(m map[string]int) {
	m["one"] = 2
}
