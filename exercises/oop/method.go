package main

import "fmt"

type triangel struct {
	size int
}

type square struct {
	size int
}

func (t triangel) perimeter() int {
	return t.size * 3
}

func (t *triangel) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

func main() {
	t := triangel{size: 3}
	s := square{size: 4}

	fmt.Println("triagle size", t)
	fmt.Println("Perimeter: triagle", t.perimeter())

	t.doubleSize()
	fmt.Println("Perimeter: triangle after double size", t.perimeter())
	fmt.Println("Perimeter: sqare", s.perimeter())
}
