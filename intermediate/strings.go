package intermediate

import (
	"fmt"
// 	"regexp"
	"strings"
// 	"unicode/utf8"
)

func main(){
	// str := "Hello, World!"
	// fmt.Println("Length of string:", len(str))
	// fmt.Println("First character:", str[0])
	// fmt.Println("Substring (0-5):", str[0:5])
	
	// str1 := "Hello"
	// str2 := "World"
	// concatenated := str1 + ", " + str2 + "!"
	// fmt.Println("Concatenated string:", concatenated)
	
	// countries := []string{"USA", "Canada", "Mexico"}
	// joined :=strings.Join(countries, ", ")
	// fmt.Println("Joined countries:", joined)

	// fmt.Println((strings.Contains(str,"GO")))
	// replaced := strings.Replace(str, "World","Go", 1)
	// fmt.Println("Replaced string:", replaced)

	// strwspace := "   Hello, World!   "
	// trimmed := strings.TrimSpace(strwspace)
	// fmt.Println("Trimmed string:", trimmed)
	// fmt.Println(strings.TrimSpace(strwspace))

	// fmt.Println(strings.ToUpper(str))
	// fmt.Println(strings.ToLower(str))

	// fmt.Println(strings.Repeat("foo",3))
	// fmt.Println(strings.Count("Hello","l"))

	// fmt.Println(strings.HasPrefix("Hello","he"))
	// fmt.Println(strings.HasSuffix("Hello","lo"))

	// str5 := "hello, 123 Go! 11"
	// re := regexp.MustCompile(`\d+`)
	// matches:=re.FindAllString(str5,-1)
	// fmt.Println("Matches:", matches)

	// str6 := "hello,    world"
	// fmt.Println(utf8.RuneCountInString(str6))

	//String Builder
	var builder strings.Builder

	//write some strings to the builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	//convert builder to string
	result := builder.String()
	fmt.Println("Built string:", result)

	//using writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are You")
	result = builder.String()
	fmt.Println("Updated string:", result)

	//reset the builder
	builder.Reset()
	builder.WriteString("New String")
	fmt.Println("After reset:", builder.String())
}