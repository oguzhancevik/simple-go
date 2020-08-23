package main

import (
	"fmt"
	"time"
)

func main() {

	goLaunchDate := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)
	// Print Date as string
	fmt.Printf("Go launch date: %s\n", goLaunchDate)

	now := time.Now()
	fmt.Printf("time now is %s\n", now)

	fmt.Println("The Current Month: ", now.Month())
	fmt.Println("The Current Day: ", now.Day())
	fmt.Println("The Current Weekday: ", now.Weekday())

	// Add 1 day
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println("tomorrow: ", tomorrow)

	fmt.Printf("Tomorrow is: %v, %v-%v-%v \n", tomorrow.Weekday(), tomorrow.Day(), tomorrow.Month(), tomorrow.Year())

	myDateFormat := "1/2/06"
	fmt.Println("Tomorrow is: ", tomorrow.Format(myDateFormat))

}
