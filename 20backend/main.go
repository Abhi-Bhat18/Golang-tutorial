package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This is about backend in GO!")
	PerformGetRequest()
	PerformPost("http://localhost:3000/")
}

func PerformPost(url string) {
	//fake json payload
	requestBody := strings.NewReader(`{"name":"Abhishek Bhat","age":21,"email":"abhishekbhat344@gmail.com"}`)
	res,err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data,err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println("Response body:", content)
}

func PerformGetRequest() {
	const myURL = "http://localhost:3000"
	res, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("")
	fmt.Println("Response status:", res.StatusCode) //we can do if res.StatusCode == 200 then do something
	fmt.Println("Response content length:", res.ContentLength)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Response body:", string(data))
	var responseString strings.Builder
	responseString.Write(data)
	content := responseString.String()
	fmt.Println("Response body:", content)
}
