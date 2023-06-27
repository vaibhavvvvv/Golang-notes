package main

import "fmt"

func main() {
	fmt.Println("Learning about FUnctions : ")

	sum, str := adder(2, 3)
	fmt.Println(str, sum)

	proSum := proAdder(1, 2, 4, 5, 6)
	fmt.Println("Sum of any number of intergers passed is : ", proSum)
}

// we can return multiple values of any datatype with the help of signatures
func adder(x int, y int) (int, string) {
	return x + y, " Sum of given nos is : "
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
