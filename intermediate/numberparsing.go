package intermediate

import (
	"fmt"
	"strconv"
)

func main(){
	numStr := "12345"
	num,err := strconv.Atoi(numStr)
	if err != nil{
		fmt.Println("Error parsing the value: ",err)
	}
	fmt.Println("parsed Integer: ",num)
	fmt.Println("Parsed Integer: ", num+1)

	numistr,err := strconv.ParseInt(numStr,10,32)
	if err != nil{
		fmt.Println("Error Parsed the value ",err)
	}
	fmt.Println("Parsed Integer: ",numistr)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr,64)
	if err != nil{
		fmt.Println("Error Parsing value: ",err)
	}
	fmt.Printf("Parsed Float: %.2f\n", floatval)

	binaryStr := "1010"
	decimal,err := strconv.ParseInt(binaryStr,2,64)
	if err != nil{
		fmt.Println("Error Parsing binary value: ",err)
	}
	fmt.Println("Parsed bianry value to decimal: ", decimal)

	hexStr := "FF"
	hexdecimal,err := strconv.ParseInt(hexStr,16,64)
	if err != nil{
		fmt.Println("Error Parsing Hex value: ",err)
	}
	fmt.Println("Parsed Hex value to decimal: ", hexdecimal)

	invalidNum := "456abc"
	invalidParse,err := strconv.Atoi(invalidNum)
	if err != nil{
		fmt.Println("Error Parsing value: ",err)
	}
	fmt.Println("Parsed invalid number: ", invalidParse)


}
