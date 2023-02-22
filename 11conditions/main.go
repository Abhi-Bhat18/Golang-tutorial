package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular Users"
	}else if loginCount > 10 {
		result = "Important User"
	}else {
		result = "Mid user"
	}
	fmt.Println(result)

	//on the go initialization in web request
	if num := 3; num <10 {
		fmt.Println("num in less than 10")
	}else{
		fmt.Println("num in not less than 10")
	}
}