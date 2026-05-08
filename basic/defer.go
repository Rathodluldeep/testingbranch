package basic

import "fmt"

func main(){
	process(10)
}

func process(i int){
	defer fmt.Println("deffer i value ",i)
	defer fmt.Println("deffer 1statement")
	defer fmt.Println("deffer 2 statement")
	defer fmt.Println("deffer 3 statement")
	i++
	fmt.Println("normal statement")
	fmt.Println("Value of i ",i)
}
