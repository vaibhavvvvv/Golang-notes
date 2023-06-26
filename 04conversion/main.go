package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter ur rating : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating , ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to you rating : ", numRating+1)
	}

}
