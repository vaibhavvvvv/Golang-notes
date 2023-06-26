package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome t0 user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter yhr rating for our Pizza : ")

	//comma ok || error ok syntax. like try-catch syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of this rating is %T", input)

}
