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

func (s square) perimeter() int {
	return s.size * 4
}

func main() {
	t := triangel{size: 3}
	s := square{size: 4}

	fmt.Println("Perimeter: triagle", t.perimeter())
	fmt.Println("Perimeter: sqare", s.perimeter())
}
