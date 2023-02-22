package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.geeksforgeeks.org/constants-go-language/"

func main() {
	fmt.Println("Welcome to handling urls in GO!")
	fmt.Println("My URL:",myUrl)

	result, _ := url.Parse(myUrl)
	fmt.Println("SCHEME:",result.Scheme) 
	fmt.Println("HOST:",result.Host) 
	fmt.Println("Path:",result.Path) 
	fmt.Println("PORT:",result.Port()) 
	fmt.Println("Raw query:",result.RawQuery) 
	
	//Getting query params from the URL
	qprams := result.Query()
	fmt.Printf("Query params: type of %T",qprams)
	fmt.Println("Query params:",qprams)

	//Creating a URL
	partsOfURL := &url.URL{ //always pass a referece of the url.URL
		Scheme: "https",
		Host: "madhuhomestaysirsi.com",
		Path: "rooms",
	}

	originalURL := partsOfURL.String()
	fmt.Println("Original URL:",originalURL)

}
