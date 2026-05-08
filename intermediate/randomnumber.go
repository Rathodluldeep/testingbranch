package inetrmediate

import (
	"fmt"
	"math/rand"
	// "time"
)

func main(){
	//  fmt.Println(rand.Intn(100)+1)
	//  fmt.Println(rand.Intn(6)+5)

	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(val.Intn(6)+5)
	// fmt.Println(rand.Float64()) //between 0.0 and 1.0

	for {
		//show the menu
		fmt.Println("welcome to the dice game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Println("enter Your choice(1 to 2 ): ")

		var choice int 
		_,err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2){
			fmt.Println("Invalid choice,plewase enter 1 or 2.")
			continue
		}
		if choice == 2{
			fmt.Println("Thanks for playing! goodbye")
			break
		}
		die1 := rand.Intn(6)+1
		die2 := rand.Intn(6)+1

		//show the results
		fmt.Println("You rolled a %d and a %d.\n",die1,die2)
		fmt.Println("total: ",die1+die2)

		//Ask of the users want to roll again
		fmt.Print("Do you wnat to roll again:? (y/n) ")
		var rollAgain string
		_,err = fmt.Scan(&rollAgain)
		if err != nil ||(rollAgain != "y" && rollAgain != "n"){
			fmt.Println("Invalid input,assuming no")
			break
		}
		if rollAgain == "n"{
			fmt.Println("Thanks for playing ! Goodbye")
			break
		}
	}

}
