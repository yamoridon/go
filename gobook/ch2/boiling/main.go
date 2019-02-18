package main

import (
	"fmt"
	"math"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	f = math.Sqrt(32.0)
	fmt.Printf("boiling point = %g°F or %g °C\n", f, c)
}
