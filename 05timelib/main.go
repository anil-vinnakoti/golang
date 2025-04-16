package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// format
	formattedPresentTime := presentTime.Format("01-02-200- 15:04:05 Monday")

	fmt.Println(formattedPresentTime)

	// create date
	createDate := time.Date(1998, time.March, 28, 10, 25, 0, 0, time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("02-01-2006 15:04:05 Monday"))
}
