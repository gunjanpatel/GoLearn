package main

import (
	"fmt"
)

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")

		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("Given number is negative")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
