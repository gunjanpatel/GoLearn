package main

import "fmt"

func main() {
	firstName := "Gunjan"
	lastName := "Patel"

	fmt.Println("Before:", firstName, lastName)
	update(&firstName, lastName)
	fmt.Println("After:", firstName, lastName)
}

func update(firstName *string, lastName string) {
	*firstName = "Prital"
	// Not going to change as it's not pointer
	lastName = "Baldha"
}
