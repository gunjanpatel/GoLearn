package main

import (
	"fmt"
	"math/rand"
	"time"
)

func infiniteBreak() {
	var num int32
	rand.New(rand.NewSource(time.Now().Unix()))

	// Infinite or unconditional loop - WOW
	for {
		fmt.Print("Writing inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}

func sumOfExcludingMod5Numbers() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
}

// વિલંબિત - Deffered
// મુલતવી રાખવું - Defer
func deferred() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}

func main() {
	var num int64
	// Deprecated
	//rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().Unix()))

	// Using like while loop
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}
}
