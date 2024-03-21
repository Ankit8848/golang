package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time World")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// to format the date, day and time use this format only.
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	oldDate := time.Date(2018, time.December, 7, 20, 20, 0, 0, time.UTC)
	fmt.Println(oldDate)
	fmt.Println(oldDate.Format("01-02-2006 Monday"))

}
