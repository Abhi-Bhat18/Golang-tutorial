package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in Go!")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createDate := time.Date(2001,time.June,19,2,3,20,0,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("2006-01-02 15:04:05 Monday"))

}