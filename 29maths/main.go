package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Maths in Golang")

	// var numberOne int = 2
	// var numberTwo float64 = 4.5
	// fmt.Println("the sum is :", numberOne+int(numberTwo))
	// float converted in int loses precision of value

	//random number
	// rand.Seed(time.Now().UnixNano()) // we seed rand with time bcoz it always changes hence random
	// fmt.Println(rand.Intn(5) + 1)    //always exclusive i.e range 0-4. we get 1-5 by +1

	//random number by cryptography. much accurate
	randNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randNum)
}
