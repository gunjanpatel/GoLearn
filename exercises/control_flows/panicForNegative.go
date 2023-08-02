package main

import (
	"fmt"
)

func main() {
	val := 0
	defer func() {
		handler := recover()

		if handler != nil {
			fmt.Println(handler)
		}
	}()

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}

func find(val int) bool {
	for {
		fmt.Println("Enter number: ")

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
