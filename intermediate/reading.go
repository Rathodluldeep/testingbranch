package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main (){
	file,err:=os.Open("output.txt")
	if err != nil{
		fmt.Println("error opening file ",err)
		return
	}
	defer func(){
		fmt.Println("closing open file")
		file.Close()
	} ()
	fmt.Println("File opened successfully: ")

// 	//Read data from file
// 	data:= make([]byte,1024)
// 	_,err = file.Read(data)
// 	if err != nil{
// 		fmt.Println("error reading from file ",err)
// 		return
// 	}
// 	fmt.Println("Data read from file: ",string(data))
// }
	scanner := bufio.NewScanner(file)

	//read line by line
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println("Read line: ",line)
	}
	err = scanner.Err(); 
	if err != nil{
		fmt.Println("error reading file ",err)
		return
	}
}