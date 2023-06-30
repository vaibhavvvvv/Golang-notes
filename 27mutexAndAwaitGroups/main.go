package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition.")

	score := []int{0}

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(4)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("1-routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("2-routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("3-routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("4-routine")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}

// when multiple goroutines access the memory simultanously for read and write, it can create problems
// hence, memory is locked for a while will operation is completed and then unlocked for others to access

// Mutex holds a lock for both reads and writes, whereas RwLock treats reads and writes differently,
// allowing for multiple read locks to be taken in parallel but requiring exclusive access for write locks.
