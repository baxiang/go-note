package main
import "fmt"

func main() {
	defer fmt.Println("ByeBye")
	fmt.Println("Hello")
}