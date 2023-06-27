package main

import "fmt"

func main() {
	fmt.Println("Learning If Else in Golang :")

	count := 13

	if count < 10 {
		fmt.Println("low count")
	} else if count > 10 {
		fmt.Println("high count")
	} else {
		fmt.Println("right count")
	}

	if marks := 55; marks > 40 {
		fmt.Println("passed")
	} else {
		fmt.Println("fail")
	}
}
