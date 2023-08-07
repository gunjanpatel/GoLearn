package main

import "fmt"

type triangel struct {
	size int
}

type coloredTriangle struct {
	triangel
	color string
}

// capital "S" means "Public" method
func (t *triangel) SetSize(size int) {
	t.size = size
}

func (t triangel) perimeter() int {
	return t.size * 3
}

// Overloading
func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := coloredTriangle{triangel{3}, "red"}

	fmt.Println("Permiter(original)", t.triangel.perimeter())

	t.SetSize(5)
	fmt.Println("Permiter(new size)", t.triangel.perimeter())
	fmt.Println("Permiter(colored)", t.perimeter())
	fmt.Println("Size", t.size)
	fmt.Println("Color", t.color)
}
