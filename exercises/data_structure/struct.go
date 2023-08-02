package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee := Employee{FirstName: "Gunjan", LastName: "Patel", Address: "Jyllinge"}

	fmt.Println("Before", employee)

	employeeCopy := &employee

	employeeCopy.FirstName = "Prital"

	fmt.Println("After", employee)
	fmt.Println(employeeCopy)
}
