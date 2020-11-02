package main

import (
	"fmt"
	"time"
)

func stop(w chan<- string) {
	fmt.Println("stop!!")
	close(w)
}

func emit(wc chan<- string, dc <-chan bool) {
	defer stop(wc)
	word := []string{"tu", "wa", "ga", "pat"}

	var count int
	t := time.NewTimer(time.Second * 2)

	for {
		select {
		case wc <- word[count]:
			count++
			if len(word) == count {
				count = 0
			}
		case <-dc:
			return
		case <-t.C:
			return
		}
	}
}

func main() {
	wordChan := make(chan string)
	doneChan := make(chan bool)

	go emit(wordChan, doneChan)

	for v := range wordChan {
		fmt.Println(v)
	}
}
