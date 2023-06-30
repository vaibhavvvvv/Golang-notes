package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //these are pointers

func main() {
	// go greeter("hello")
	// greeter("world")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1) //adds to goroutines. keeps track of them
	}

	wg.Wait() // waits until everyone returns
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)

// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() // informs when all is done.

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOps in endpoint")
	} else {
		fmt.Printf("%d status code of %s\n", res.StatusCode, endpoint)
	}
}
