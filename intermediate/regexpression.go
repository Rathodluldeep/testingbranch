package intermediate

import (
	"fmt"
	"regexp"
)

func main(){
	fmt.Println("He said, \"I am Great\"")
	fmt.Println(`he said, "I am great"`)

	//complie a regex pattern  to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-z]{2,}`)

	//test string
	email1:="user1@email.com"
	email2:= "invalid_email"

	//match
	fmt.Println("email1: ",re.MatchString(email1))
	fmt.Println("email2: ",re.MatchString(email2))

	// Capturing Groups
	// complie a regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	//Test string
	date := "2024-07-30"

	//find all submatch
	submatch := re.FindStringSubmatch(date)

	fmt.Println(submatch)
	fmt.Println(submatch[0])
	fmt.Println(submatch[1])
	fmt.Println(submatch[2])
	fmt.Println(submatch[3])

	//Source string
	str := "Hello World"

	re =regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str,"*")
	fmt.Println(result)

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)

	//Test string
	text := "Golanf is going great"

	//Match
	fmt.Println("Match: ",re.MatchString(text))
}