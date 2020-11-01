package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup
var counter int

const total = 10

func foo() {
	for i := 0; i < total; i++ {
		counter += i
		runtime.Gosched()
		go func() {
			fmt.Println("Inside Foo loop :", counter, " -> i:", i)
		}()
		waitGroup.Done()
	}
	fmt.Println("counter 1:", counter)
}

func main() {
	waitGroup.Add(total)
	go foo()
	waitGroup.Wait()
}
