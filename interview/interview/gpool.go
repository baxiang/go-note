package main

import "fmt"

func maxString(s string)int{
	m :=make(map[byte]int)
	for i:=0;i<len(s);i++{
		m[s[i]]++
	}
	fmt.Println(m)
	var res int
	var isDouble bool
	for _,v:=range m{
		r := v%2
		if r==0{
			res +=v
		}else if !isDouble{
			res+=1
			isDouble= true
		}
	}
	return res
}

func longStr(s string)string{
	dp := make([][]bool,len(s))
	for i:=range s{
		dp[i]=make([]bool,len(s))
	}
	res :=""
	for i:=len(s)-1;i>=0;i-=1{
		for j:=i;j<len(s);j+=1{
			dp[i][j]=s[i]==s[j]&&(j-i<3||dp[i+1][j-1])
			if dp[i][j]&&j-i+1>len(res){
				res =s[i:j+1]
			}
		}
	}
	return res
}


func main() {
	fmt.Println(longStr("babad"))
}
