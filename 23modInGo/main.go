package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in GO language")
	greater()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//running as a server
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greater() {
	println("Hello there!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to go lang series</h1>"))
}
