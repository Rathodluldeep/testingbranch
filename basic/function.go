package basic

import "fmt"

func main(){
	//func <name>(parameters list) returnType{
	//codeblock
	//return value
	//}
	fmt.Println(add(2,3))
	sum:= add(2,4)
	fmt.Println(sum)
	greet()

	operation:= add
	result := operation(3,5)
	fmt.Println(result)
}

func add(a,b int)int{
	return a+b
}

var greet = func(){
	fmt.Println("Hello anonymous function")
}