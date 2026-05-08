package intermediate

import (
	"fmt"
	"os"
)


func checkError(err error){
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
}

func main (){
	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	// checkError(os.MkdirAll("subdir1", 0755))
	
	// defer os.RemoveAll("subdir1")
	// checkError(os.WriteFile("subdir1/file",[]byte(""), 0755)) 
	// checkError(os.MkdirAll("subdir1/parent/child", 0755))
	// os.WriteFile("subdir/parent/child/file",[]byte(" "), 0755)

	result,err := os.ReadDir("subdir1/parent")
	checkError(err)
	for _,entry := range result{
		fmt.Println(entry.Name(),entry.IsDir(),entry.Type())
	}
	
	checkError(os.Chdir("subdir1/parent/child"))

	result,err = os.ReadDir(".")
	checkError(err)
	for _,entry := range result{
		fmt.Println(entry.Name(),entry.IsDir(),entry.Type())
	}
	checkError(os.Chdir("../../.."))
	dir,err := os.Getwd()
	checkError(err)
	fmt.Println(dir)
}