package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum, substract, mul := calc(os.Args[1], os.Args[2])
	fmt.Printf("Sum: %d\nSustract: %d\nMultiplication: %d\n", sum, substract, mul)

	only_sum, _, _ := calc(os.Args[1], os.Args[2])
	fmt.Printf("Only Sum: %d \n", only_sum)
}

func sum_basic(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	return int1 + int2
}

func sum(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	result = int1 + int2

	return
}

func calc(number1 string, number2 string) (sum int, substract int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	sum = int1 + int2
	substract = int1 - int2
	mul = int1 * int2

	return
}
