package main

import "fmt"

func main() {
	fmt.Println("Welcome to learning Pointers")

	// var ptr *int
	// fmt.Println("Value of ptr pointer is", ptr) //<nil>

	myNumber := 23
	var ptr = &myNumber // reference passed with &. contains address of myNumber

	fmt.Println("Value of ptr is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr) //23.
	fmt.Println("Value of *ptr = *ptr + 2 is : ", *ptr+2)

}

// why pointers exist? : sometimes when variables are to be passed, their copy is passed
//instead which creates irregularities. To avoid this, we use Pointers.
// Pointers are direct reference to memory address.

// *ptr will always contain exact value which is used to pass
// &variable contains address of a variable.
