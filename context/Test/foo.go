package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	ioutil.WriteFile("./tt.txt", []byte("WriteString writes a ## string"), 0666)
	fileObj, _ := os.Open("./tt.txt")
	Rd := bufio.NewReader(fileObj)
	content, _ := Rd.ReadSlice('#')
	fmt.Println(string(content))
}
