package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Current time is")

	currentTime := time.Now()

	fmt.Println(currentTime.Format("01/02/2006 Monday") )

	myBirthday := time.Date(1997, time.March, 14, 7, 30, 43, 53, time.UTC)

	fmt.Println(myBirthday)
}