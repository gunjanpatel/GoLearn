package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gunjanpatel/calculator"
	"rsc.io/quote"
)

func main() {
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])

	total := calculator.Sum(number1, number2)

	fmt.Println("Total", total)
	fmt.Println(calculator.Version)
	fmt.Println(quote.Glass())
}
