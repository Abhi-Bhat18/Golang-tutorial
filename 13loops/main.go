package main

import "fmt"

func main() {
	fmt.Println("loops in GO!")

	// days := []string{"Sunday","Monday","Tuesday","Wednesday","Thursday"}

	//type 1
	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	//Traversing in slice or array
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index,day := range days {
	// 	fmt.Printf("Day %v: %v\n",index,day)
	// }

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2{
			goto scoreMatch
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("value is: ",rougeValue)
		rougeValue++
	}

	scoreMatch:
		fmt.Println("score matched ")
}