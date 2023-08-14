package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkApi(endpoint string) {
	_, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("ERROR: %s is down\n", endpoint)
		return
	}

	fmt.Printf("SUCCESS: %s is up and running\n", endpoint)
}

// "Do not communicate by sharing memory; instead, share memory by communicating."
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

	for _, endpoint := range apis {
		go checkApi(endpoint)
	}

	elapsed := time.Since(start)

	fmt.Printf("Done! It took %v seconds!\n", elapsed)

	// Done! It took 0s seconds!
}
