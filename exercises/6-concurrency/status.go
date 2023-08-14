package main

import (
	"fmt"
	"net/http"
	"time"
)

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
		_, err := http.Get(endpoint)

		if err != nil {
			fmt.Printf("ERROR: %s is down\n", endpoint)
			continue
		}

		fmt.Printf("SUCCESS: %s is up and running\n", endpoint)
	}

	elapsed := time.Since(start)

	fmt.Printf("Done! It took %v seconds!\n", elapsed)

	// Output without concurrency - check status-concucrrancy.go file to see result with concurrency.
	// SUCCESS: https://management.azure.com is up and running
	// SUCCESS: https://dev.azure.com is up and running
	// SUCCESS: https://api.github.com is up and running
	// SUCCESS: https://outlook.office.com/ is up and running
	// ERROR: https://api.somewhereintheinternet.com/ is down
	// SUCCESS: https://graph.microsoft.com is up and running
	// Done! It took 2.0626153s seconds!
}
