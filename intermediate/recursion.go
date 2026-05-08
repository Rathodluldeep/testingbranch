package intermediate

import "fmt"

func main(){
	fmt.Println(factorial(5))
	fmt.Println(factorial((10)))

	fmt.Println(sumOfDigit(9))
	fmt.Println(sumOfDigit(123))
	fmt.Println(sumOfDigit(12345))
}

func factorial(n int)int{
	//base case :factorial of 0 and 1
	if n ==0{
		return 1
	}
	//Recursion case: factorial of n is n * factorial (n-1)
	return n* factorial(n - 1)
	//n*(n - 1) * (n - 2)*factorial ( n - 3) *  factorial(0)
}
func sumOfDigit(n int)int{
	//Base case
	if n < 10{
		return n
	}
	return n%10 + sumOfDigit(n/10)
}