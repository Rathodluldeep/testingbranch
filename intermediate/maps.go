package main

import "fmt"

func main() {
	//var mapVCariable map[key Type]valueType\
	//mapVariable=make(map[key Type]valueType)
	//using  a map literal
	//key1:value1
	//key2=value2

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key"] = 9
	fmt.Println(myMap)
}
