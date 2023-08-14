package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	size float64
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func (s Square) Area() float64 {
	return s.size * s.size
}

type Triangle struct {
	size float64
}

func (t Triangle) Perimeter() float64 {
	return t.size * 3
}

func (t Triangle) Area() float64 {
	return t.size * t.size
}

func showInfo(s Shape) {
	fmt.Println("Your shape is:", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
