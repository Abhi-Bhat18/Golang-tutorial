package main

import "fmt"


const Logintoken string = "gibbershih"

func main() {
	var username string = "Abhishek"
	fmt.Println(username)
	fmt.Printf("Username is of type: %T \n", username)

	var isAdmin bool = true
	fmt.Println(isAdmin)
	fmt.Printf("isAdmin is of type: %T \n", isAdmin)

	var age uint8 = 21
	fmt.Println(age)
	fmt.Printf("age is of type: %T \n", age)

	var cgpa float32 = 8.24
	fmt.Println(cgpa)
	fmt.Printf("age is of type: %T \n", cgpa)

	// default values and some aliases

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("anotherVar is of type: %T \n", anotherVar)

	// implicit type - no need to specify type of the variable
	var website = "madhuhomestaysirsi.com"
	fmt.Println(website)

	// no var style (no need of var keyword)
	porjectName := "GoLang"
	fmt.Println("\nproject name is declared using the no var style of declaration")
	fmt.Println(porjectName)

	// constants
	fmt.Println("\nConstants")
	fmt.Println(Logintoken)

}