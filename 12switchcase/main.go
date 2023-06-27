package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Learning  Switch Case")

	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1")
	case 2:
		fmt.Println("dice value is 2")
		fallthrough //unconditianally executes next case
	case 3:
		fmt.Println("dice value is 3")
		fallthrough
	case 4:
		fmt.Println("dice value is 4")
	case 5:
		fmt.Println("dice value is 5")
	case 6:
		fmt.Println("dice value is 6")
	}

}
