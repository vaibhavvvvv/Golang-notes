package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Learning web req verbs")

	performGetRequest()
}

func performGetRequest() {

	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	//fmt.Println(string(content))
	// another way to convert byte data into string is : using strings module and Builder method.

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(responseString.String()) //various mathods available for strings builder
}

// Output :
// Learning web req verbs
// Status Code : 200
// Content length :  43
// {"message":"Hello from learnCodeonline.in"}
