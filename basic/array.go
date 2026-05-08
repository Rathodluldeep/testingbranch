package basic

import "fmt"

func main (){
	//. var arrayName [size]elementType
	var numbers[5]int
	fmt.Println(numbers)

	numbers[4]=20
	numbers[0]=9
	fmt.Println(numbers)

	fruits:=[4]string{"apple","banana","orange","Grapes"}

	fmt.Println(fruits)

	orignalArray:=[3]int{1,2,3}
	//copiedArray:=orignalArray
	var copiedArray *[3]int

	copiedArray=&orignalArray

	copiedArray[0]=100
	fmt.Println("orignalArray: ",orignalArray)
	//fmt.Println("copiedArray: ",copiedArray)
	fmt.Println("copiedArray: ",*copiedArray)

	for i :=0;i<len(numbers);i++{
		fmt.Println("Element at index,",i,":",numbers[i])
	}
	for index,value := range numbers{
		fmt.Printf("Index: %d,Value:%d\n",index,value)
	}
	fmt.Println(len(numbers))

	array1:=[3]int{1,2,3}
	array2:=[3]int{10,2,3}

	fmt.Println("array is equal to Array2:",array1==array2)

	a,b:=someFunction()
	fmt.Println(a,b)

	c := 2
	_= c

}
 

func someFunction()(int, int){
	return 1,2
}