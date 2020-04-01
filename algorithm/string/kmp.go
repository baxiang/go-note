package main

func main() {
	
}



func next(findStr string)(next map[int]int){

	next = map[int]int{0:0}
	var k int
	for i :=1;i<len(findStr);i++{
		for k >0 && findStr[k]!=findStr[i]{
			k = next[k-1]
		}
		if findStr[k] == findStr[i]{
			k++
		}
		next[i] = k
	}
	return
}