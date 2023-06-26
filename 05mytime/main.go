package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study time in golang")
	presentTime := time.Now()
	fmt.Println(presentTime)                                      //ugly output
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05")) //formats beautifully but...compulsary parameter to pass

	createdtime := time.Date(2023, time.July, 16, 12, 12, 12, 12, time.UTC)
	fmt.Println(createdtime.Format("01-02-2006 Monday 15:04:05"))
}

//...........................................................................................

// to build file for any OS just type in terminal :  GOOS="windows" go build
// corresponding build file will be created.
// 1) type : go env
// 2) find : GOOS="your os name"
