package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Learning modules in GO")

	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello there, welcome to golang.")
}

// func to read request and write a response
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go. Home.</h1>")) //byte data format for web
}

//
// go list : lists the dependant modules
// go list all : lists all dependant modules
// go list -m all : lists major used modules and packages
// go list -m -versions github.com/gorilla/mux : lists all the available versions of mux
//
// go mod why github.com/gorilla/mux : asks why am I dependant on this package?
// go mod graph :  pulls up graph of all dependencies

// go mod edit -go 1.16 : to change version of go
// go mod edit -module 1.16: to change module's name
//
//go mod vendor : constructs a directory named vendor that contains copies of packages. similar to node_modules

//go run main.go : brings packages to run the program from web/cache but not vendor
// go run -mod=vendor main.go : checks if there is vendor. If yes the run using it otherwise run using web/cache
