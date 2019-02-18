package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	x := 9.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
