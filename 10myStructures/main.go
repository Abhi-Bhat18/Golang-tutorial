package main

import "fmt"

func main() {
	// no inheritance in golang; No super or parent concepts in GO
	fmt.Println("This is about structures in GO")

	//Structures are the user defined data structures
	abhishek := User{"abhishek","abhishekbhat344@gmail.com",true,16}
	
	// fmt.Println(abhishek)
	fmt.Printf("User name is: %v and email: %v ",abhishek.Name,abhishek.Email)

	// fmt.Printf("hitesh details are:%+v\n",abhishek)
}

//Defining a structre
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
