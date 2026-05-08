package intemediate

import (
	"encoding/json"
	"fmt"
	"log"
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
	
	jsondata1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsondata1))

	jsonData1 := `{"first_name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`
	
	//Unmarshalling
	var employeeFromJson Employee
	err=json.Unmarshal([]byte(jsonData1),&employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Age + 5: ",employeeFromJson.Age + 5)
	fmt.Println("City: ",employeeFromJson.Address.City)

	listOfCityStates := []Address{
		{City: "New York", State: "NY"},
		{City: "Los Angeles", State: "CA"},
		{City: "Chicago", State: "IL"},
		{City: "Houston", State: "TX"},
		{City: "Phoenix", State: "AZ"},
	}
	fmt.Println(listOfCityStates)
	jsonList,err := json.Marshal(listOfCityStates)
	if err != nil {
		log.Fatalln("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonList))
	//handaling unknown Json Structures
	jsonData2:= `{"first_name": "Alice", "age": 28, "address": {"city": "San Francisco", "state": "CA"}, "unknown_field": "some value"}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(data)
}

type Employee struct{
	FirstName string `json:"first_name"`
	EmpID int `json:"emp_id"`
	Age int `json:"age,omitempty"`	
	Address Address `json:"address"`
}