package intermediate

import "fmt"

func main(){
	// sequence := adder()

	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// sequence2 := adder()
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())

	substractor := func() func(int) int{
		countdown := 99
		return func (x int) int{
			countdown -= x
			return countdown
		}
	}()

	// Using the closure substraction

	fmt.Println(substractor(1))
	fmt.Println(substractor(2))
	fmt.Println(substractor(3))

}

func adder() func() int{
	i := 0
	fmt.Println("Previous value of i : ",i)
	return  func() int{
		i++ 
		fmt.Println("adder 1 to i")
		return i
	}
}