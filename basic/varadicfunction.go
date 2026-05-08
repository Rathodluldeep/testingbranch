package basic

import "fmt"

func main(){
	//..eclipsis
	// func functionName(param1 type1 ,param2 type2 ,param3 ...type3)returnType{
	// 	function bodyconst
	// }
	//fmt.Println("sum of 1,2,3:", sum(1,2,3))

	// statement,total := sum("the sum of 1,2,3 is ", 1,2,3,4,5,6)
	// fmt.Println(statement,total)

	sequence,total := sum(1,20,30,40,50,60)
	fmt.Println("Sequence: ",sequence,"total: ",total)
	sequence1,total1 := sum(2,20,30,40,50,60)
	fmt.Println("Sequence: ",sequence1,"total: ",total1)
	
	numbers:=[]int{1,2,3,4,5,9}
	sequence3, total3 :=sum(3,numbers...)
	fmt.Println( "sequence: ",sequence3,"total: ",total3)
	
}

func sum(sequence int ,nums ...int)(int, int){
		total:=0
		for _,v := range nums{
			total += v
		}
		return sequence, total
	}