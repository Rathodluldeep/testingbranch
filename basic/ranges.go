package basic

import "fmt"

func main(){
	message:="hello world"

	for i,v := range message{
		//fmt.Println(i,v)
		fmt.Printf("Index: %d,Rune:%c \n", i,v)
	}
}