package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{
		Circle: Circle{
			Point:  Point{8, 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)

	w.X = 42

	fmt.Printf("%#v\n", w)
}
