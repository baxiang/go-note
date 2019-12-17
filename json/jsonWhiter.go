package main

import (
	"encoding/json"
	"fmt"
	"os"
)
type StuGender int

const (
	MALE = 1 << iota
	FEMALE
)

type Student struct {
	Id     int64     `json:"id"`
	Name   string    `json:"name"`
	Gender StuGender `json:"gender"`
	Exams  []exam    `json:"exams"`
}
type exam struct {
	Lesson string `json:"lesson"`
	Score  int    `json:"score"`
}
func main() {

	stu := Student{Id: 101, Name: "小明", Gender: MALE, Exams: []exam{{Lesson: "Chinese", Score: 98}, {Lesson: "Math", Score: 100}}}
	// 存储文件
	f, err := os.Create("./student.json")
	if err != nil {
		fmt.Println(err)
	}
	if err = json.NewEncoder(f).Encode(stu);err!=nil{
		fmt.Println(err)
	}
	// 读取文件
	file, err := os.Open("./student.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var s Student
	if err = json.NewDecoder(file).Decode(&s);err==nil{
		fmt.Println(s)
	}else {
		fmt.Println(err)
	}



}
