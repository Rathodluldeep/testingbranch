package basic
import (
	"fmt"
	foo "net/http"
)

func main (){
	fmt.Println("go standard library")

	resp, err :=foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err !=nil{
		fmt.Println("Error",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Http Response Status: ",resp.Status)
}