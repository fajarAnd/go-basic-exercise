package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("DALAM -> Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Goroutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("runtime goroutine:", runtime.NumGoroutine())
	fmt.Println("Final Counter:", counter)
}
