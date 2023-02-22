package main

import "fmt"

//function starts with func keyword and followed by function name

func main() {
	fmt.Println("Functions in GO!")
	// result := adder(3,5)
	// fmt.Println(result)

	proResult := proAdder(3,23,12,32,12,23)
	fmt.Println(proResult	)

}

func adder(valOne int,valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for index := range values{
		total += values[index]
	}
	return total
}

func greeter (){
	fmt.Println("Namaste from Abhishek")
}