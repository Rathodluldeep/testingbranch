package intermediate	

import (
	"fmt"

	"go.uber.org/zap"
)

func main (){
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Print("Error in intializing Zap logger")
	}
	defer logger.Sync()

	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// logger.Info("user logged in",zap.String("username","John Doe"),zap.String("method","GET"))

	//log with fields
	logger.With(
		zap.String("user", "john"),
		zap.String("method", "GET"),
	).Info("User logged in")

}