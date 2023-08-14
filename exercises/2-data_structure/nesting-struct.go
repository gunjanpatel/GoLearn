package main

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Person
	ManagerId int
}

// emplouee.PersonInfo.FirstName
// type Employee struct {
// 	PersonInfo Person
// 	ManagerId int
// }

type Contractor struct {
	Person
	ContractorId int
}

func main() {
	employe := Employee{
		Person: Person{
			FirstName: "Gunjan",
		},
		ManagerId: 1,
	}

	contractor := Contractor{
		Person: Person{
			LastName: "Patel",
			Address:  "Jyllinge",
		},
		ContractorId: 3,
	}

	fmt.Println("Employee", employe)
	fmt.Println("Employee Firstname", employe.FirstName)
	fmt.Println("Employee Manager ID", employe.ManagerId)
	fmt.Println("Contractor", contractor)
	fmt.Println("Contractor Firstname", contractor.FirstName)
	fmt.Println("Contractor Lastname", contractor.LastName)
	fmt.Println("Contractor Address", contractor.Address)
}
