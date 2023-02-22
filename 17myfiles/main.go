package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is about file system in GO!")
	content := "This needs to go in a file - www.geekydev.in"

	//Creating a file 
	file,err := os.Create("./file1.txt")
	checkNilErr(err)
	
	len,err := io.WriteString(file,content)
	checkNilErr(err)

	fmt.Println("length is: ",len)
	defer file.Close()
	readFile("./file1.txt")

}

//Reading a file
func readFile(filename string){
	dataByte,err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file in \n",string(dataByte))
}


//checking error 
func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}