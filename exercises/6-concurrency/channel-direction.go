package main

import "fmt"

func read(ch <-chan string) {
	fmt.Println("Receiving message:", <-ch)
}

func send(ch chan<- string, message string) {
	fmt.Println("Sending message", message)
	ch <- message
}

func main() {
	ch := make(chan string, 1)
	send(ch, "Hello")
	read(ch)
}
