package main

import "fmt"

func main() {
	scores := map[string]int{"chinese": 102, "math": 136, "english": 115}

	for k, v := range scores {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	for k := range scores {
		fmt.Printf("key: %s\n", k)
	}

	for _, v := range scores {
		fmt.Printf("value: %d\n", v)
	}
}
