package intermediate
import "fmt"

func main(){
	num := 42
	fmt.Printf("%05d\n", num) // Output: 00042
// 5 digits, left-padded with zeros
	message := "hello"
	fmt.Printf("|%10s|\n", message) // Output: |hello|
	fmt.Printf("|%-10s|\n", message) // Output: |hello     |

	message1:="Hello \n World!"
	message2:="Goodbye \n World!"

	fmt.Println(message1)
	fmt.Println(message2)
	//sqlQuer:=`SELECT * FROM users WHERE name = 'John' AND age > 30;`
	//fmt.Println(sqlQuer)
}