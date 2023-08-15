// Switch statement but for channels
package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done Processing!!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done Replicating!!"
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go process(ch1)
	go replicate(ch2)

	start := time.Now()

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process, time.Since(start))
		case replicate := <-ch2:
			fmt.Println(replicate, time.Since(start))
		}
	}
}
