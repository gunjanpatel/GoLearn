package main

import "fmt"

func main() {
	val := 0
	for {
		fmt.Print("Enter number:")
		fmt.Scanf("%d", val)

		if val != 0 && val <= 2 {
			fmt.Println("Enter the number greater than 2")
		} else if val > 0 {
			fmt.Println(fibonacci(val))
			break
		}
	}
}

func fibonacci(number int) []int {
	output := []int{1}
	sum := 1
	for i := 1; i < number; i++ {
		output = append(output, sum)
		sum = output[i] + output[i-1]
	}

	return output
}
