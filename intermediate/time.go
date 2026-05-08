package intermediate

import (
	"fmt"
	"time"
)

func main(){
	//Current local time
	fmt.Println(time.Now())
	//Specific time
	specificTime := time.Date(2024,time.July,30,12,0,0,0,time.UTC)
	fmt.Println("Specific Time: ",specificTime)

	//Parsed time
	parsedTime, _ := time.Parse("2006-01-02","2020-05-01")  //Mon Jan 2 15:01:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02","20-05-01")      //Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)

	//formating time
	t := time.Now()
	fmt.Println("Formatted toijme: ",t.Format("Monday 06-01-02 15-04-05"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())
	fmt.Println("Rounded Time: ",t.Round(time.Hour))

	// loc,_:=time.LoadLocation("Asia/Kolkata")
	// t=time.Date(2024, time.July, 8,14,16,40,00, time.UTC)

	// //Convert this to specific time zone
	// tLocal := t.In(loc)

	// //Perform rounding
	// roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)

	// fmt.Println("Original Time (UTC): ", t)
	// fmt.Println("Original Time (Local): ", tLocal)
	// fmt.Println("Rounded Time (roundedTime): ", roundedTime)
	// fmt.Println("Rounded  Time (Local): ", roundedTimeLocal)

	fmt.Println("Truncated Time: ",t.Truncate(time.Hour))
}
