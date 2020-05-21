package main

import (
	"encoding/json"
	"strings"
	"fmt"
)



func CloudmallInterview1(numbers []int) []int {
	//posIndex :=0
	// 先查找+
	slow := 0
	for numbers[slow]>0{
		slow++ // 找到第一个非正数的位置
	}
	fast :=slow+1//
	for fast<len(numbers){
		if numbers[fast]>0 {
			numbers[slow],numbers[fast] = numbers[fast],numbers[slow]
			slow =slow+1
			for numbers[slow]>0{
				slow++ // 找到非正数的位置
			}
			if slow>fast {
				fast=slow+1
			}
		}else {
			fast++
		}

	}
	for numbers[slow]==0{
		slow++ // 找到第一个非正数的位置
	}
	fast =slow+1
	for fast<len(numbers){
		if numbers[fast]==0 {
			numbers[slow],numbers[fast] = numbers[fast],numbers[slow]
		}
		for numbers[slow]==0{
			slow++ // 找到第一个非正数的位置
		}
		fast++
	}
	return numbers
}
func CloudmallInterview2(numbers []int) []int {
	//posIndex :=0
	// 先查找+
	left := 0
	right :=len(numbers)-1
	for left<right{
		for numbers[right]<=0{
			right--
		}
		if numbers[left]<=0 {
			numbers[left],numbers[right]= numbers[right],numbers[left]
		}
		left++
	}
	right =len(numbers)-1
	for left<right{
		for left<right&&numbers[right]!=0{
			right--
		}
		if numbers[left]<0 {
			numbers[left],numbers[right]= numbers[right],numbers[left]
		}
		left++
	}
	return numbers
}



func CloudmallInterview3(revList map[string]string) (jsonStr string, err error) {

	var res []interface{}
	resMap  := make(map[string]interface{})
	for k ,v:=range revList{
		l := strings.Split(v,".")
		if len(l)==0 {
			resMap[v]=k
		}else {
			for i,_:=range l{
				for j:=0;j<i;j++ {
					var preMap map[string]interface{}
                  if r,ok:= resMap[l[j]];ok{
					  preMap = r.(map[string]interface{})
				  }else {
					  if j==0 {
						  resMap[l[j]]= k
					  }else {
						  //preMap[l[j-1]] =make(map[string]interface{}){""}
					  }

				  }
				}
			}
		}
	}
	b ,err := json.Marshal(res)

	return string(b), err
}

func isExestMap(m map[string]interface{},k []string){
	for i:=0;i<len(k);i++{

	}
}


func main() {
	//fmt.Println(CloudmallInterview2([]int{6, 4, -3, 0, 5, -2, -1, 0, 1, -9}))
	//fmt.Println(CloudmallInterview2([]int{-6, -4, -3, 0, 5, 2, 1, 0, 1, 9}))

	//revList := map[string]string{"1": "bar", "2": "foo.bar", "3": "foo.foo", "4": "baz.cloudmall.com", "5": "baz.cloudmall.ai"}


	revList1 := map[string]string{ "2": "foo.bar", "3": "foo.foo"}
	res,_:= CloudmallInterview3(revList1)
	fmt.Println(res)

}
