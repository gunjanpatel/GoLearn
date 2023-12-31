package main

import (
	"fmt"
	"time"
)

func fib(number float64, ch chan<- string) {
	x, y := 1.0, 1.0

	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func main() {
	start := time.Now()

	size := 15
	ch := make(chan string, size)

	for i := 1; i < size; i++ {
		go fib(float64(i), ch)
		select {
		case m := <-ch:
			fmt.Printf(m)
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Done!! It took %v seconds!\n", elapsed)
}
