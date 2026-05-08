package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main (){
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("Home: ", home)
	fmt.Println("User: ", user)

	err:=os.Setenv("fruit","Apple")
	if err != nil {
		fmt.Println("Error setting environment variable",err)
	}
	fmt.Println("Fruit: ", os.Getenv("fruit"))

	// for _,env := range os.Environ(){
	// 	Kvpair := strings.SplitN(env,"=",2)
	// 	fmt.Println(Kvpair[0])
	// }
	//"a=b=c=d"
	//n=1 returns "a=b=c=d"
	//n=2 returns "a" and "b=c=d"
	//n=3 returns "a","b","c=d"
	//n=4 returns "a","b","c","d"

	err=os.Unsetenv("fruit")
	if err != nil {
		fmt.Println("Error unsetting environment variable",err)
		return
	}
	fmt.Println("Fruit: ", os.Getenv("fruit"))

	str := "a=b=c=d=e"
	fmt.Println(strings.SplitN(str,"=", -1))
	fmt.Println(strings.SplitN(str,"=", 0))
	fmt.Println(strings.SplitN(str,"=", 1))
	fmt.Println(strings.SplitN(str,"=", 2))
	fmt.Println(strings.SplitN(str,"=", 3))
	fmt.Println(strings.SplitN(str,"=", 4))
	fmt.Println(strings.SplitN(str,"=", 5))
}
