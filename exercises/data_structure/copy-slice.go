package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	// Changing slice1 will also change slice2 if we don't use copy
	slice2 := letters[1:4]

	// So let's use copy to avoid problem
	slice3 := make([]string, 3)
	copy(slice3, letters[1:4])

	fmt.Println("Before Slice2", slice2)
	fmt.Println("Before Slice3", slice3)

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("After Slice2", slice2)
	fmt.Println("After Slice3", slice3)
}
