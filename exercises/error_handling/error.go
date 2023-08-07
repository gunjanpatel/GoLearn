package main

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrEmployeeNotFound = errors.New("Employee not found")

func main() {
	employee, err := getInformation(1001)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(employee)
	}
}

func getInformation(id int) (*Employee, error) {

	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(id)

		if !errors.Is(err, ErrEmployeeNotFound) {
			return employee, nil
		}

		fmt.Println("Server is not responding, retrying...")

		// Retry after 2 second
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*Employee, error) {
	if id != 1000 {
		return nil, ErrEmployeeNotFound
	}
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
