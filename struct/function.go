package main

import "fmt"




type StudentFun struct {
	name string
	age  int
	score  float32
}
func (s StudentFun) updateAge(age int){
	s.age = age
}
func (s *StudentFun) updateScore(score float32){
	s.score = score
}

func main() {
	s := StudentFun{"li", 21, 100}
	fmt.Println(fmt.Sprintf("before:%v", s))
	s.updateAge(18)
	s.updateScore(60)
	fmt.Println(fmt.Sprintf("update:%v", s))


}
