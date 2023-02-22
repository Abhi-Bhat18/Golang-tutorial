package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods in GO!")
	fmt.Println("Methods are the functions written in Structures")
	
	abhishek := User{"Abhishek","abhishekbhat344@gmail.com",true,21}
	abhishek.ChangeRole()

	fmt.Println("original user details are: ",abhishek)
}


type User struct{
	Name string
	Email string
	IsAdmin bool 
	Age int
}

func (u User) ChangeRole(){
	fmt.Println("user Status before: ", u.IsAdmin)
	u.IsAdmin = false
	fmt.Println("Role of the user after: ", u.IsAdmin)
}