package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Learning web req verbs")

	//performGetRequest()
	//performPostJsonRequest()
	performPostFormRequest()
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

func performPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//to create fake json data/payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "golang",
			"price" : "0",
			"platform" : "avc.com"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content)) // {"coursename":"golang","price":"0","platform":"avc.com"}

	defer response.Body.Close()
}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// to create fake form data .....url module
	data := url.Values{}
	data.Add("firstname", "vaibhav")
	data.Add("lastname", "gadhave")
	data.Add("email", "abcd@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content)) // {"email":"abcd@gmail.com","firstname":"vaibhav","lastname":"gadhave"}

	defer response.Body.Close()
}
