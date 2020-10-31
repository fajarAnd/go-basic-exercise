package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo1() {
	arr := [10]int{}

	for i, _ := range arr {
		fmt.Println("loop Arr:", i)
	}
	fmt.Println("Foo1 done:", time.Now())
	wg.Done()
}

func foo2() {
	slice := make([]string, 50, 50)
	for i, _ := range slice {
		fmt.Println("Loop Slice:", i)
	}
	fmt.Println("Foo2 done:", time.Now())
	wg.Done()
}
func main() {
	fmt.Println("number goroutine before:", runtime.NumGoroutine())

	wg.Add(2)
	go foo1()
	go foo2()

	fmt.Println("number goroutine after:", runtime.NumGoroutine())
	wg.Wait()
}
