package main

import (
	"fmt"
	"regexp"
	"time"
)

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

func weekday() {
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())
}

func validate() {
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}
}

func fall() {
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	// This condition will be ignored and code block will be executed
	// Eventhogh this condition doesn't match code block will still be executed
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	// This condition will be ignored and code block will be executed
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}

func main() {
	region, continent := location("Copenhagen")
	fmt.Printf("Gunjan works in %s, %s\n", region, continent)
}
