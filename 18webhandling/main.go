package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("This is abuot get")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error is ", err)
		panic(err)
	}
	data, err := ioutil.ReadAll(response.Body)

	content := string(data)
	fmt.Println(content)
	defer response.Body.Close()

}
