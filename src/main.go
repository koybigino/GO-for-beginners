package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { // send the data to the channel only
		i := <-ch
		fmt.Println(i)
		// ch <- 27
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // receive the data to the channel only
		ch <- 42
		// fmt.Println(<-ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
