package main

import "fmt"

func main() {
	fmt.Println("Learn defer :::")

	defer fmt.Println("Hello")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//output
// 4
// 3
// 2
// 1
// 0
// three
// two
// one
// Hello

// defer :
// --any expression written infront of defer will get exexcuted at the end of function i.e.
//   exactly before "}"
// --for multiple defers, it follows LIFO order. it executes defer in reverse order.
