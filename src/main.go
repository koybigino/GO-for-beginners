package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // takes 50 int numbers
	wg.Add(2)
	go func(ch <-chan int) { // send the data to the channel only
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // receive the data to the channel only
		ch <- 42
		ch <- 27
		close(ch) // close the channel
		wg.Done()
	}(ch)

	wg.Wait()
}
