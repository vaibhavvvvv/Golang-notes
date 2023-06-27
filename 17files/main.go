package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Learning file handling in golang.")

	content := "this is a content to be entered in a file."

	//os module to create the file
	file, err := os.Create("./myfile.txt")
	checkoutNilError(err)

	// io module to write or manipulate the file
	// io.WriteString returns length of file.
	length, err := io.WriteString(file, content)
	checkoutNilError(err)

	fmt.Println("Length of file is :", length)

	// compulsary to close file at the end of function
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	// ioutil is used to read file
	// file data is sent in the form of bytes.
	// convert it to string by using string()
	byteData, err := ioutil.ReadFile(filename)
	checkoutNilError(err)

	fmt.Println("Text inside the file is :\n", string(byteData))
}

// necessary to check errors after each file operation
func checkoutNilError(err error) {
	if err != nil {
		panic(err) // panic stops the execution and throws the error if occured.
	}
}
