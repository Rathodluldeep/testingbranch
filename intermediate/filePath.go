package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)
	

func main (){
	relativePath := "./data/file.txt"
	absolutePath := "/home/user/data/file.txt"
	//join paths using filepath package
	joinedPath := filepath.Join("downloads", "file.zip")
	fmt.Println("Joined Path: ",joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path: ",normalizedPath)
	

	dir,file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("Directory: ",dir)
	fmt.Println("File: ",file)

	fmt.Println(filepath.Base("/home/user/docs/"))

	fmt.Println("Is relativePath variable absolute",filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute",filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))

	fmt.Println(strings.TrimSuffix(file,filepath.Ext(file)))


	rel , err:= filepath.Rel("/home/user", "/home/user/docs/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Relative path from /home/user to /home/user/docs/file.txt: ",rel)

	rel , err= filepath.Rel("/home/user", "/home/user/docs/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Relative path from /home/user to /home/user/docs/file.txt: ",rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error getting absolute path: ",err)
	}else{
	fmt.Println("Absolute path: ",absPath)
	}
}