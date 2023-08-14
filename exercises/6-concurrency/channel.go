package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(endpoint string, ch chan string) {
	_, err := http.Get(endpoint)

	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down", endpoint)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running", endpoint)
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

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	// We can close the channel as now we know nothing more is going to add into it
	// close(ch)
	// Without closeing if we execute this then programme will continue running
	// fmt.Println(<-ch)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed)
}
