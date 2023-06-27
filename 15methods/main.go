package main

import "fmt"

func main() {
	fmt.Println("Learning Methods in Golang")

	//there is no inheritance in golang; Nor super or parent as it makes coding complex.
	//similar to a class

	Vaibhav := User{"Vaibhav", "vaibhavng7@gmail.com", true, 20}
	fmt.Println("Info about Vaibhav is ; ", Vaibhav)
	fmt.Printf("Info about Vaibhav is %+v\n; ", Vaibhav)
	fmt.Printf("Vaibhav's age is %v and email is %v \n", Vaibhav.Age, Vaibhav.Email)

	Vaibhav.getStatus()
	Vaibhav.NewMail()                                              // abc@gmail.com
	fmt.Printf("Vaibhav's original email is %v \n", Vaibhav.Email) //vaibhavng7@gmail.com
	// For a method, a copy of object is passed and not the original object.
	// Hence, methods don't affect values of original object unless pointer is used.

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Status of user is : ", u.Status)
}

func (u User) NewMail() {
	newmail := "abc@gmail.com"
	fmt.Println("New mail is : ", newmail)
}

//methods are just the functions in classes/structs.
