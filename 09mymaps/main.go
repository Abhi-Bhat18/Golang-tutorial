package main

import "fmt"

func main() {

	fmt.Println("Maps in go lang ")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["TS"] = "TypeScript"
	languages["GO"] = "GoLang"

	fmt.Println("Programming values: ", languages)

	fmt.Println("Primary language: ", languages["JS"])

	//delete method 
	delete(languages,"GO")

	fmt.Println("languages: ", languages)

	//looping in go lang

	for key,value := range languages{
		fmt.Printf("for key %v, value is %v\n",key,value)
	}
}