package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("word.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if len(buf) > 0 {
			fmt.Println(string(buf))
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//offset = offset + n
	}
}
