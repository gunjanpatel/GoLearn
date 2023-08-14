package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(endpoint string, ch chan string) {
	_, err := http.Get(endpoint)

	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!!", endpoint)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is working fine.", endpoint)
}

func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	size := 2
	ch := make(chan string, size)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)

	fmt.Printf("Done! It took %v", elapsed)
}
