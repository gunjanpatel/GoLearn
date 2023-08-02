package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Copenhagen", "Aaruhus", "Aalborg":
		region, continent = "Denmark", "Europe"
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Gujarat":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("Copenhagen")
	fmt.Printf("Gunjan works in %s, %s\n", region, continent)
}
