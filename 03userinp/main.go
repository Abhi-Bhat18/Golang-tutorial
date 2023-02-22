package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program!"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	//comma ok syntax or err ok syntax --> It takes a input or an error
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)

	fmt.Println("Enter your age: ")
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age is", age)

	
}