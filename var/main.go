package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 1.0, 2
	c := math.Min(a, float64(b))
	fmt.Println("minimum value = ", c)
}