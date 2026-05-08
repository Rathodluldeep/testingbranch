package intermediate

import "github.com/sirupsen/logrus"

func main (){
	log :=logrus.New()

	//set log level
	log.SetLevel(logrus.InfoLevel)

	//set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	//logging example
	log.Info("this is info message")
	log.Warn("this is warning message")
	log.Error("this is error message")
	
	//log with fields
	log.WithFields(logrus.Fields{
		"user": "john",
		"method": "GET",
	}).Info("User logged in")


}