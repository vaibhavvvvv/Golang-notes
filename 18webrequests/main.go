package main

import (
	"fmt"
	"io/ioutil"
	"net/http" // net/http is used for handling web requests
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Learning to handle web requests in Golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response is type of : %T \n", response) //*http.Response . pointer used to access and manipulate original value

	defer response.Body.Close() //caller's responsibilty to close the connection.

	byteData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("content of response body is ", string(byteData))
}
