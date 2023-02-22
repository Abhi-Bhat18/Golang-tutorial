package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to slices in Go!")

	//declaring a array 
	// var myarr = [5]int{1,2,3,4,5} // we need to specify the size of the array
	// fmt.Println("Array is: ", myarr)

	//declaring a slice
	// var friends = []string{"Raj","Rani","Raju","Rani"} // no need of defining the size of the slice
	// fmt.Println("friends(slice): ", friends)

	//appending the data to slice
	// friends = append(friends, "Rocky","Adi","Ganesh")
	// fmt.Println("Added friends: ", friends)

	//slicing in slice
	// fmt.Println("old friends: ", append(friends[4:]))

	// highScore := make([]int,4) //by default it is an array 
	// highScore[0] = 234;
	// highScore[1] = 224;
	// highScore[2] = 23;
	// highScore[3] = 343;
	// highScore = append(highScore, 4323,234323,23323)
	// fmt.Println("scores: ", highScore)

	// //sorting in slices
	// sort.Ints(highScore)
	// fmt.Println("sorted Highscores: ", highScore)

	//remove the value from the slices based on the index 
	skills := []string{"reactJs","swift","nodejs","mongodb","expressjs"}
	fmt.Println("skills: ",skills)

	//removing the skills

	skills = append(skills[:2],skills[3:]...)

	fmt.Println(skills)

}