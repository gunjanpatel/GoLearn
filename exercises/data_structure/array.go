package main

import "fmt"

func main() {
	// Size of Array is ALWAYS fix
	var manual [3]int

	manual[1] = 10
	// Default for last elmenet is 0 as data type default
	fmt.Println(manual[0], manual[1], manual[len(manual)-1])

	// Dynamic length assignment
	cities := [...]string{"Copenhagen", "Ahmedabad"}

	fmt.Println("Cities", cities)

	// Setting value only to the last element or specific element
	n := [...]int{99: -1}

	fmt.Println("First", n[0])
	fmt.Println("Last", n[len(n)-1])
	fmt.Println("Total", len(n))

	// twoDdimentional array
	var twoD [3][5]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)

	// Three dimentional
	var threeD [2][3][2]int

	fmt.Println(threeD)
}
