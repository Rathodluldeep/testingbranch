package main

import (
	"encoding/json"
	"fmt"
	"log"
)


type Person struct {
	FirstName string `json:"first_name"` //"db:"first_name"`
	LastName string `json:"last_name,omitempty"`
	Age int `json:"-"`        //"age,omitempty" ("-" means ignore this field during JSON marshaling)
}

func main (){
	// person :=Person{FirstName: "John", LastName: "Doe", Age: 30}
	person :=Person{FirstName: "John", LastName: "", Age: 20}
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

}
