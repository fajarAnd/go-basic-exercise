package main

import (
	"fmt"
	"sync"
	"time"
)

func populate(p chan int) {
	for i := 0; i < 20; i++ {
		p <- i
		fmt.Println("Populate: ", i, "---->", time.Now().Nanosecond())
	}
	close(p)
}

func fanInOut(i, o chan int) {
	var wg sync.WaitGroup
	for v := range i {
		wg.Add(1)
		go func(v2 int) {
			o <- v2
			fmt.Println("FanInOut: ", v2, "--->", time.Now().Nanosecond())
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(o)
}

func main() {
	x := make(chan int)
	y := make(chan int)

	go populate(x)
	go fanInOut(x, y)

	for v := range y {
		fmt.Println("---- bekerja:", v, "--->", time.Now().Nanosecond())
	}
	fmt.Println("Cabut", time.Now())
}
