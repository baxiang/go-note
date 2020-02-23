package main

import "fmt"

func judgeScore(score int) string {
	var result string
	if score >= 90 {
		result = "优"
	} else if score >= 80 {
		result = "良"
	} else if score >= 70 {
		result = "中"
	} else if score >= 60 {
		result = "及格"
	} else {
		result = "继续努力"
	}
	return result
}

func main() {
	fmt.Println(judgeScore(75))
	fmt.Println(judgeScore(40))
	fmt.Println(judgeScore(95))
	fmt.Println(judgeScore(63))
	fmt.Println(judgeScore(86))
}
