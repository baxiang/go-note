package main

import (

	"fmt"

	"sort"

	"strings"

)


type StuScore struct {

	name  string

	score int

}


type StuScores []StuScore


func (s StuScores) Len() int {

	return len(s)

}


func (s StuScores) Less(x, y int) bool {

	return s[x].score < s[y].score

}

func (s StuScores) Swap(x, y int) {

	s[x], s[y] = s[y], s[x]

}


type SortByName struct {

	StuScores

}


func (s SortByName) Less(x, y int) bool {

	if strings.Compare(s.StuScores[x].name, s.StuScores[y].name) < 1 {

		return true

	}

	return false

}


type SortByAgeDESC struct {

	StuScores

}


func (s SortByAgeDESC) Less(x, y int) bool {

	return s.StuScores[x].score > s.StuScores[y].score

}


func main() {

	stus := StuScores{{"name2", 95}, {"name1", 75}, {"name5", 86}, {"name4", 60}, {"name3", 100}}

	fmt.Println("排序前------------------")

	for _, v := range stus {

		fmt.Println(v.name, ":", v.score)

	}

	fmt.Println("按照分数升序排序------------------")

	sort.Sort(stus)

	for _, v := range stus {

		fmt.Println(v.name, ":", v.score)

	}

	fmt.Println("按照姓名排序------------------")

	sort.Sort(SortByName{stus})

	for _, v := range stus {

		fmt.Println(v.name, ":", v.score)

	}

	fmt.Println("按照分数降序排序------------------")

	sort.Sort(SortByAgeDESC{stus})

	for _, v := range stus {

		fmt.Println(v.name, ":", v.score)

	}

}

