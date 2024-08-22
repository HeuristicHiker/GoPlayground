package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	t := triangle{
		base:   0.78,
		height: 4,
	}
	s := square{
		sideLength: 2,
	}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	a := 0.5 * t.base * t.height
	return a
}

func (s square) getArea() float64 {
	a := s.sideLength * s.sideLength
	return a
}

func printArea(s shape) {
	fmt.Println("Area is: ", s.getArea())
}
