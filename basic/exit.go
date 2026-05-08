package basic

import (
	"fmt"
	"os"
)

func main(){
	defer fmt.Println("Deferred statement")

	fmt.Println("Starting the main function")
	//Exit with status code of 1
	os.Exit(1)
	//this will never
	fmt.Println("End of main function")
}
