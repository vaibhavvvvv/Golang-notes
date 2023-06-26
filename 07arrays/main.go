package main

import "fmt"

func main() {
	fmt.Println("Lets learn arrays today.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Watermelon"

	fmt.Println("Fruit list is ", fruitList)                // [Apple Banana  Watermelon]. note extra space
	fmt.Println("length of Fruit list is ", len(fruitList)) // 4

	var veggieList = [3]string{"Tomato", "Potato", "onion"}
	fmt.Println("Veggies list is ", veggieList) //[Tomato Potato onion]

}

//arrays are not used much in golanf unlike other languages.
