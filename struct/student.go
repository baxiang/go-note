package main
import "fmt"
type Person struct {
	name string
	age  int
}
type Student struct {
	Person // 匿名字段
	score  float32
}
func main() {
	m := Person{"ming", 25}
	fmt.Println(fmt.Sprintf("%v", m))
	h := Person{"wang", 18}
	fmt.Println(fmt.Sprintf("%v", h))
	s := Student{Person{"li", 21}, 100}
	fmt.Println(fmt.Sprintf("%v", s))
}
