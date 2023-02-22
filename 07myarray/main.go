package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go!")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Orange"
	fruitlist[2] = "Banana"
	fruitlist[3] = "Grapes"
	
	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("length of the array is: ",len(fruitlist))

	var veglist = [3]string{"Potato","Tomato","Onion"}
	fmt.Println("Veg list is: ", veglist)
}