package main

import (
	"fmt"
)
func main() {
	fmt.Print("Welcome to pointers in Go!\r")
	fmt.Println("Pointers are used to share memory address of a variable and they are nothing but direct reference to the memory location of a variable.")
	
	myAge := 21
	var agePtr *int = &myAge  //referecne operator is &
	fmt.Println("value of the pointer is: ", agePtr)
	fmt.Println("pointer holding the value: ", *agePtr) //dereference operator is *

	*agePtr = *agePtr + 2
	fmt.Println("value of the age is now: ", myAge)
}
