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
	fmt.Printf("%s %s is %d years old\n", firstName, lastName, age)

	// Checking default values
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println("Following are the default values of each data types:")
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)
}
