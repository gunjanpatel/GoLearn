package main

import "fmt"

func main() {
	// This won't work as we must have to init map with make to allow adding in it. Add doesn't work on "nil" map. However read and delete works on nil map. GO doesn't panic there.
	// studentsAge := map[string]int
	studentsAge := make(map[string]int)

	studentsAge["gunjan"] = 32
	studentsAge["yuval"] = 4

	delete(studentsAge, "gunjan")

	// GO doesn't panic as prital is not exist in the map
	delete(studentsAge, "prital")

	fmt.Println(studentsAge)

	// No panic here as GO default to type default
	fmt.Println("Prital's age is", studentsAge["prital"])

	// Using second argument to know that value is set or not
	age, exist := studentsAge["prital"]
	if exist {
		fmt.Println("Prital's age is", age)
	} else {
		fmt.Println("Prital's age couldn't be found")
	}

	// Looping throgh map element
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
