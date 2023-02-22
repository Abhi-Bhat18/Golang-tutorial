package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello this is a json example")
	// EncodeJson()
	// DecodeJson()
	CommonJson()
}

type User struct {
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
}

func EncodeJson() {
	abhi := User{"abhishek", 30, "abhibhat344@gmail.com", "ahbi1234","1234567890"}
	userJson,err := json.MarshalIndent(abhi, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", userJson)
}


func DecodeJson () {
	jsonDataFromWeb := []byte(`{
        "fullName": "abhishek",
        "age": 30,
        "email": "abhibhat344@gmail.com",
        "password": "ahbi1234",
        "contact": "1234567890"}`)
	var abhi User 
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &abhi)
		fmt.Printf("%#v\n", abhi)
	} else {
		fmt.Println("Json is not valid")
	}

}

func CommonJson (){
	var myData map[string]interface{}
	json.Unmarshal([]byte(`{"name":"abhishek","age":30}`), &myData)

	// fmt.Printf("%#v\n", myData)
	for key,val := range myData {
		fmt.Printf("%s: %v\n", key, val)
	}
}