package main

import (
	"fmt"
)

type shape interface {
	calculateArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (s square) calculateArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) calculateArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("The area is ", s.calculateArea())
}

func main() {
	s := square{5.0}
	t := triangle{}
	t.height = 10
	t.base = 3
	printArea(s)
	printArea(t)

}
