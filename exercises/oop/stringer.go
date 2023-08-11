package main

import "fmt"

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%s is from %s", p.Name, p.Country)
}

func main() {

	p1 := Person{Name: "Gunjan", Country: "Denmark"}
	p2 := Person{Name: "Hardik", Country: "Australlia"}

	fmt.Printf("%s\n%s\n", p1, p2)
}
