package basic

import "fmt"

func main(){
	//panic(interface{})
	//Example of valid input

	process(10)

	//example of invalid input

	process(-3)
}

func process(input int){
	defer fmt.Println("Deffered 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be non-negative number")
		//fmt.Println("After panic")

		//defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input: ", input)
}