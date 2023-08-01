package main

import "fmt"

func main() {
	// Ways to declare variables
	// var firstName, lastName string
	// var (
	// 	firstName, lastName string
	// 	age int
	// )
	// var (
	// 	firstName = "Gunjan"
	// 	// firstName string = "Gunjan"
	// 	lastName  = "Patel"
	// 	//lastName string = "Patel"
	// 	age = 32
	// 	//age int = 32
	// )
	firstName, lastName := "Gunjan", "Patel"
	age := 32
	fmt.Printf("%s %s is %d years old", firstName, lastName, age)
}
