package main

import "fmt"

func main() {
	fmt.Println("Learning Maps : ")

	//creating maps
	var languages = make(map[string]string)

	languages["Py"] = "Python"
	languages["Js"] = "Javascript"
	languages["Ts"] = "typescript"

	fmt.Println("Languages in map are : ", languages)  //map[Js:Javascript Py:Python Ts:typescript]
	fmt.Println("Js is short for : ", languages["Js"]) // Javascript

	//deleting elements from map
	delete(languages, "Py")
	fmt.Println("updated languages : ", languages)

	//loops.   smimlar to for each loop in js
	for key, value := range languages {
		fmt.Printf("for key %v , value is %v.\n", key, value)
	}
}
