package main

import "fmt"

func main() {
	// Error: cannot use -10 (untyped int constant) as uint value in variable declaration (overflows)
	var integer uint = -10
	fmt.Println(integer)
}
