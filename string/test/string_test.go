package test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkForm(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fmt.Sprintf("%d", i)
	}
}
func BenchmarkItoa(t *testing.B) {
	for i := 0; i < t.N; i++ {
		strconv.Itoa(i)
	}
}
