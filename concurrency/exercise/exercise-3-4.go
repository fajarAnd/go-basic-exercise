package main

import (
	"sync"
	"fmt"
)

var waitGroup sync.WaitGroup

const total = 10

func foo() {
	var counter int
	var mutex sync.Mutex
	for i := 0; i < total; i++ {
		go func() {
			mutex.Lock()
			counter += i
			//runtime.Gosched()
			fmt.Println("Inside Foo loop :", counter, " -> i:", i)
			mutex.Unlock()
			waitGroup.Done()
		}()
	}
	fmt.Println("counter 1:", counter)
}

func main() {
	waitGroup.Add(total)
	foo()
	waitGroup.Wait()
}
