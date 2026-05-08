package basic

import (
	"fmt"
)

func main() {
	// var slicesName[]ElementType

	// var numbers []int
	// var numbers1 =[]int{1,2,3}
	// numbers2 := []int{9,8,7}

	// slice := make([]int,5)

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("SliceCopy, ", sliceCopy)

	// var nilSlice []int
	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index of slice", slice1[3])

	// slice[3]=50
	// fmt.Println("Element at index of 3 slice",slice[3])

	// for slices.Equal(slice1,sliceCopy){
	// 	fmt.Println("slice1 is equal to sliceCopy")
	// }
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j:=0;j<innerLen; j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println(twoD)

	//slice[low:high]
	slice2:= slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("the capacity of cslice2 is ",cap(slice2))
	fmt.Println("the len of slice2 is ",len(slice2))
}
