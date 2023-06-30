package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and Deadlocks")

	mych := make(chan int, 2) // increase size of buffer to handle multiple inputs
	wg := &sync.WaitGroup{}

	//results in deadlock : fatal error: all goroutines are asleep - deadlock!
	// mych <- 5
	// fmt.Println(<-mych)

	wg.Add(2)
	//send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 99
		close(ch) //sends output as 0 if it has not input.i.e. channel close

		ch <- 5
		ch <- 6 // handled with bufer
		wg.Done()
	}(mych, wg)

	//recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		//to handle confusion between closed channel output(0) and input(0)
		val, isChannelOpen := <-ch
		fmt.Println("Is channel open :", isChannelOpen)
		fmt.Println("input in channel: ", val)

		fmt.Println(<-ch)
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
