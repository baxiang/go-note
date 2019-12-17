package main

import (
	"encoding/json"
	"fmt"
)

//type StuGender int
//
//const (
//	MALE = 1 << iota
//	FEMALE
//)
//
//type Student struct {
//	Id     int64     `json:"id"`
//	Name   string    `json:"name"`
//	Gender StuGender `json:"gender"`
//	Exams  []exam    `json:"exams"`
//}
//type exam struct {
//	Lesson string `json:"lesson"`
//	Score  int    `json:"score"`
//}



func main() {
	ming := Student{Id: 101, Name: "小明", Gender: MALE, Exams: []exam{{Lesson: "Chinese", Score: 98}, {Lesson: "Math", Score: 100}}}
	data, _ := json.Marshal(ming)

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err == nil {
		fmt.Println(m)
	}

	for k,v :=range m{
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "是string", vv)
		case bool:
			fmt.Println(k, "是bool", vv)
		case float64:
			fmt.Println(k, "float64", vv)
		case nil:
			fmt.Println(k, "是nil", vv)
		case []interface{}:
			fmt.Println(k, "是array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "未知数据类型")
		}
	}
}
