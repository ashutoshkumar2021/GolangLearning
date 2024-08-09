package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()   //helps to take the current time
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))  // this is the fixed date, month,year and time and day we have to use always to give formate

	createDate :=time.Date(2016,time.March,30,07,07,0,0,time.UTC)   //time.Date able to create own date
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
