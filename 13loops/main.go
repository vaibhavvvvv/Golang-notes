package main

import "fmt"

func main() {
	fmt.Println("Lerning loops, goto,break, continue")

	days := []string{"monday", "tuesday", "wedneday", "thursday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for d := range days {
	// 	fmt.Println(d, days[d]) //unlike in js, for each does not return element but its index
	// }
	// 0 monday
	// 1 tuesday
	// 2 wedneday
	// 3 thursday

	for index, day := range days {
		fmt.Printf("index is %v and day is %v.\n", index, day)
	}

	//to skip index and just get value
	for _, day := range days {
		fmt.Printf("index is  and day is %v.\n", day)
	}

// 	rogueValue := 1

// label:
// 	fmt.Println("Hello label")

// 	for rogueValue < 10 {
// 		if rogueValue == 2 {
// 			goto label
// 			continue
// 		}

// 		if rogueValue == 5 {
// 			rogueValue++
// 			break
// 		}
// 		rogueValue++
// 	}

// }
