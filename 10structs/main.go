package main

import "fmt"

func main() {
	fmt.Println("Learning Structs in Golang")

	//there is no inheritance in golang; Nor super or parent as it makes coding complex.
	//similar to a class

	Vaibhav := User{"Vaibhav", "vaibhavng7@gmail.com", true, 20}
	fmt.Println("Info about Vaibhav is ; ", Vaibhav)
	fmt.Printf("Info about Vaibhav is %+v\n; ", Vaibhav)
	fmt.Printf("Vaibhav's age is %v and email is %v ", Vaibhav.Age, Vaibhav.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
