package main

import (
	"encoding/json"
	"fmt"
)

// we can replace the keywords appearing in JSON using the below way.
// don't leave any whitespaces. strict format.
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Password string   `json:"-"`
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Learn : creating JSON data in golang")

	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	onlineCourses := []course{
		{"Reactjs", 200, "abc123", "youtube", []string{"web-dev", "js"}},
		{"Nodejs", 300, "ddd123", "udemy", []string{"app-dev", "js"}},
		{"Golang", 400, "vvv123", "coursera", nil}}

	// package above data as JSON data
	//	MarshalIndent is used to implement JSON. 1st: variable, 2nd: prefix to add on each indent, 3rd: how to indent?
	finalJson, err := json.MarshalIndent(onlineCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	// %s used to display json
	fmt.Printf("%s\n", finalJson)
}

// decoding JSON data that's received
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Reactjs",
		"Price": 200,
		"website": "youtube",
		"tags": ["web-dev","js"]
	}
	`)

	var newCourses course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &newCourses) // converts json format to bytes format into original newCourses(& and not a copy)
		fmt.Printf("%#v", newCourses)
	} else {
		fmt.Println("Json is invalid")
	}

	var myOnlineData map[string]interface{} //creating a map
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
