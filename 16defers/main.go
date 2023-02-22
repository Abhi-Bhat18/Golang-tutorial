package main

import "fmt"

func main() {
	defer fmt.Println("First in defers")
	defer fmt.Println("Second in defers")
	defer fmt.Println("Third in defers")
	fmt.Println("This is after defers in order")
	myDefers()
}

func myDefers(){
	for i:=0; i<4; i++ {
		defer fmt.Println(i)
	}
}
