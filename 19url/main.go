package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=1252344234"

func main() {
	fmt.Println("Learning how to work/handle URLs in Golnag")

	fmt.Println(myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Println("result is : ", result)

	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //lco.dev:3000
	fmt.Println(result.Port())   //3000
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=1252344234
	fmt.Println(result.Path)     //   /learn

	qparams := result.Query()
	fmt.Printf("type of query params is : %T \n", qparams) //url.Values

	for _, val := range qparams {
		fmt.Println("Params is : ", val) // [reactjs]    1252344234]

	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	newUrl := partsOfUrl.String()
	fmt.Println("New url constructed is : ", newUrl)
}
