package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lets learn Slies today")

	var fruitList = []string{"Apple", "Peaches", "Watermelon"}
	fmt.Printf("Type of fruitList is : %T \n", fruitList) // []string

	fruitList = append(fruitList, "MAngo", "Banana") //append() adds elements to slice.
	fmt.Println("Fruitlist is : ", fruitList)        //[Apple Peaches Watermelon MAngo Banana]

	// fruitList = append(fruitList[1:]) //elements from 1st position onwards are sliced.
	fruitList = append(fruitList[1:3]) //1 onwards till 3(3 not inclusive. last range never included)

	fmt.Println("Fruitlist is : ", fruitList) //[Peaches Watermelon MAngo Banana]

	highScores := make([]int, 4)

	highScores[0] = 22
	highScores[1] = 44
	highScores[2] = 66
	highScores[3] = 33

	//highScores[4] = 55  //cant add new elements. It gives errors. So, instead we do like this:

	highScores = append(highScores, 55, 77) //[22 44 66 33 55 77]
	sort.Ints(highScores)                   //[22 33 44 55 66 77]

	fmt.Println("high scores are : ", highScores)
	fmt.Println("are high scores sorted ? : ", sort.IntsAreSorted(highScores))

	//.......................................................................................
	//removing element at index position from Slice

	var courseList = []string{"java", "python", "c++", "Go", "Typescrit"}
	index := 2 //element to be removed
	courseList = append(courseList[:index], courseList[index+1:]...)
	fmt.Println("courses after removing index element are : ", courseList)
}

//SLices: Powerful and mostly used data type in Golang
// they are under the hood Arrays. Arrays when get abstraction level and some more features, they are called slices.
// We can add as many values as we want in SLices. it automaticlly expands.
