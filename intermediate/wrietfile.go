package intermediate

import (
	"fmt"
	"os"
)


func main (){
	file,err:= os.Create("output.txt")
	if err != nil{
		fmt.Println("error creating file ",err)
		return
	}
	defer file.Close()

	//write data to file
	data := []byte("Hello, this is a sample text written to the file.")
	_, err = file.Write(data)
	if err != nil{
		fmt.Println("error writing to file ",err)
		return
	}
	fmt.Println("Data written to file successfully.")

	file, err = os.Create("writeString.txt")
	if err != nil{
		fmt.Println("error creating file ",err)
		return
	}
	defer file.Close()

	_,err =file.WriteString("Hello , Go!\n")
	if err != nil{
		fmt.Println("error writing string to file ",err)
		return
	}
	fmt.Println("String written to file successfully.")
}