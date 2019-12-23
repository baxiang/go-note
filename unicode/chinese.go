package main

import (
	"fmt"
	"unicode"
)

func ChineseStr(query string) bool {
	set := []*unicode.RangeTable{unicode.Scripts["Han"],
		unicode.Number,
		unicode.Space,
		unicode.P,
		unicode.Symbol,
		unicode.C,
		unicode.Space}

	for _, r := range query {
		if !unicode.IsOneOf(set, r) {
			return false
		}
	}

	return true
}
func main() {
    fmt.Println(ChineseStr("你好世界111"))
}
