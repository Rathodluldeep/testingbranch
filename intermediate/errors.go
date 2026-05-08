package intermediate

import (
	"errors"
	"fmt"
)
// func sqrt(x float64) (float64 , error) {
// 	if x < 0 {
// 		return 0,errors.New("math: square root of negative number")
// 	}
// 	return 1, nil
// }

type customError struct {
	Message string
	code int
	err error
}
//Error return the error message.Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error: %d (message: %s) %v\n", e.code, e.Message, e.err)
}

//function that return a custom error
// func doSomething() error {
// 	return &customError{
// 		Message: "Something went wrong",
// 		code: 500,
// 	}
// }
func doSomething() error {
	err := doSomethingElse()
	if err != nil{
		return &customError{
			Message: "Failed to do something else",
			code: 500,
			err: err,
		}
	}
	return nil
}
func doSomethingElse() error {
	return errors.New("internal error")
}

func main() {

	// result,err := sqrt(16)
	// if err !=nil{
	// 	fmt.Println("Error:",err)
	// 	return
	// }
	// fmt.Println("Result: ", result)/

	err := doSomething()
	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Operation successful")
}