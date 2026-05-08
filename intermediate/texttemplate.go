package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main(){
	// //tmpl := template.New("greeting")
	// // Define a template string with placeholders
	// tmpl := template.Must(template.New("greeting").Parse("Hello, {{.Name}}! Welcome!"))
	// // Create a template and parse the string into it
	// // tmpl,err :=template.New("greeting").Parse("Hello, {{.Name}}! Welcome!")
	// // if err != nil {
	// // 	panic(err)
	// // }

	// //define data to fill the template
	// data := map[string]interface{}{
	// 	"Name":"Alice",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Define named tamplate for diffrent types of
	 templates := map[string]string{
		"welcome":"welcome, {{.name}}! We're glad you joined.",
		"notification":"{{.name}}, you have a new notification: {{.notification}}",
		"error":"Oops! An error occured: {{.errorMessage}}",
	}

	//Parse and store template
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name]= template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// show menu
		fmt.Println("\n Menu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose and option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		
		var tmpl *template.Template

		switch choice{
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name":name}

		case "2":
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{
				"name":         name,
				"notification": "You got a new message!",
			}

		case "3":
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{
				"errorMessage": "Something went wrong!",
			}

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
			continue
		}

		//render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil{
			fmt.Println("Error execute template: ", err)
		}

	}
}