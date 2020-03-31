package main

import "fmt"

// NextIndex sets ix to the lexicographically next value,
// such that for each i>0, 0 <= ix[i] < lens(i).
func NextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

func main() {
	e := [][]string{
		{"a1"},
		{"b1", "b2"},
		{"c1", "c2", "c3"},
		{"d1"},
	}
	lens := func(i int) int {
		l := len(e[i])
		if l>2 {
			return 2
		}
		return len(e[i])

	}

	for ix := make([]int, len(e)); ix[0] < lens(0); NextIndex(ix, lens) {
		var r []string
		for j, k := range ix {
			r = append(r, e[j][k])
		}
		fmt.Println(r)
	}
}

