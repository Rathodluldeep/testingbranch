package intermediate

import (
	"log"
	"os"
	
)


func main (){
	log.Println("this is a log message")
	log.SetPrefix("INFO: ")
	log.Println("this is an info message")

	// log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Llongfile)
	log.Println("this is a log message with date,time,shortfile")

	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message")
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	infoLogger1 := log.New(logFile, "INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1 := log.New(logFile, "WARN: ",log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(logFile, "ERROR: ",log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger := log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message")
	warnLogger1.Println("This is a warning message")
	errorLogger1.Println("This is an error message")
	infoLogger1.Println("This is an info message")	

}
var(
	infoLogger = log.New(os.Stdout, "INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, "WARN: ",log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ",log.Ldate|log.Ltime|log.Lshortfile)
)