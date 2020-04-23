package main

import (
	"fmt"
	"sort"
)

type strLen struct {
	Data  string
	Length int
}
type strLenList []*strLen
func (p strLenList) Len() int           { return len(p) }
func (p strLenList) Less(i, j int) bool { return p[i].Length < p[j].Length }
func (p strLenList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func sortString(s []string)[]string{
	l :=make(strLenList,0)
	for i,v:=range s{
		sl := &strLen{Data:s[i],Length:len(v)}
		l =append(l,sl)
	}
	sort.Sort(l)
	r :=make([]string,len(l))
    for _,v:=range l{
    	//fmt.Println(v.Data)
    	r =append(r,v.Data)
	}
	return r
}

func main() {
   fmt.Println(sortString([]string{"abc","ed","afgh"}))
}
