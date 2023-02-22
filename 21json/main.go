package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello this is a json example")
	EncodeJson()
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	Password  string
	IsAdmin   bool
}
type Course struct {
	Name     string   `json:"courseName"`  //this is called struct tag
	Price    int      `json:"coursePrice"` //this is called struct tag
	Platform string   `json:"Website"`
	Password string   `json:"-"` //this is called struct tag
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	// abhi := User{"abhishek", "bhatt", 30, "abhishekbhat344@gmail.com", "1234", true}

	users := []User{
		{"abhishek", "bhatt", 30, "abhishekbhat@gmail.com", "1234", true},
		{"abhi", "bhat", 23, "abhi@gmail.com", "2343", false},
	}

	geekyDev := []Course{
		{"React", 100, "Udemy", "1234", []string{"react", "js"}},
		{"Go", 200, "Udemy", "1234", []string{"go", "js"}},
	}

	//package this data as JSON data

	userJson, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}

	finalJson, err := json.MarshalIndent(geekyDev, "", "\t")

	if err != nil {
		panic(err)
	}
	// fmt.Println(string(finalJson))

	fmt.Printf("%s", userJson)

	fmt.Printf("%s", finalJson)
}
