package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"` //`db: "user_id"`
	// LastName string `json:"last_name"`
	Age int `json:"age,omitempty"`
	EmailAddress string `json:"email,omitempty"` //omitempty means if the field is empty, it will be omitted from the JSON output
	Address Address `json:"address"`
}

type Address struct{
	City string `json:"city"`
	State string `json:"state"`
}

func main (){
	person := Person{FirstName: "John", Age: 30, Address: Address{City: "New York", State: "NY"}}

	//Marshalling
	jsonData, err:=json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 25, Address: Address{City: "Los Angeles", State: "CA"}}
	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))
}