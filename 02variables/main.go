package main

import "fmt"

const LoginToken = "dfsdsdd" //variables starting with Capital letters are Public variables.

func main() {
	var username string = "vaibhav"
	fmt.Printf(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type : %T \n", smallval)

	//default value of int is 0 and that of string is null/""
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T \n", anotherVariable)

	//another way to declare variables/ implicit type/ i.e. direct declartion. here lexer identifies type based on value provided
	var name = "namaste boiii"
	fmt.Println(name)

	//no var style.    Not allowed outside methods
	numberOfPeople := 1000000
	fmt.Println(numberOfPeople)

	fmt.Println(LoginToken)
}
