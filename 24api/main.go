package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Model for course file

type Course struct {
	CourseId   string  `json:"courseId"`
	CourseName string  `json:"courseName"`
	Price      int     `json:"prince"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}


//fake DB 
var courseList = []Course{

}

//middlewares 
func IsEmpty(c *Course) bool{
	return c.CourseId == "" && c.CourseName == "" 
}

func main() {
	fmt.Println("Hello this is API in GO language")
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	//running as a server
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to go lang series</h1>"))
}
