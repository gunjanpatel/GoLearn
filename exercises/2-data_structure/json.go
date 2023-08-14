package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string `json:"omitempty"`
	Address   string `json:"address,omitempty"`
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
	// JSON encode decode of Array
	employee := []Employee{
		{
			Person: Person{
				FirstName: "Gunjan",
			},
			ManagerId: 1,
		},
		{
			Person: Person{
				FirstName: "Prital",
				LastName:  "Patel",
			},
			ManagerId: 2,
		},
	}

	jsonData, _ := json.Marshal(employee)

	fmt.Printf("Employee JSON:\t%s\n", jsonData)

	var decodeJson []Employee

	json.Unmarshal(jsonData, &decodeJson)

	fmt.Printf("Decoded Employee:\t%v\n", decodeJson)

	// JSON encode decode of Object
	contractor := Contractor{
		Person: Person{
			FirstName: "Gunjan",
		},
		ContractorId: 2,
	}

	contractorJson, _ := json.Marshal(contractor)

	var decoded Contractor

	json.Unmarshal(contractorJson, &decoded)

	fmt.Printf("Contractor JSON:\t%s\n", contractorJson)
	fmt.Printf("Decoded Contractor:\t%v\n", decoded)
}
