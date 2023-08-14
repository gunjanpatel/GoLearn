package main

import (
	"fmt"

	"github.com/gunjanpatel/creditStore"
)

func main() {
	gunjan, _ := creditStore.CreateEmployee("Gunjan", "Patel", 500)

	fmt.Println("Initial Credit:", gunjan.CheckCredits())
	fmt.Println(gunjan)
	credits, err := gunjan.AddCredits(100)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	_, err = gunjan.RemoveCredits(1000)

	if err != nil {
		fmt.Println("Error:", err)
	}

	gunjan.ChangeFirstName("Prital")

	fmt.Println(gunjan)
}
