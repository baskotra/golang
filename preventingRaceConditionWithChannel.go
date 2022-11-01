package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	channel := make(chan bool, 1)
	for i := 0; i <= 1000; i++ {
		w.Add(1)
		go increment(&w, channel)
	}
	w.Wait()
	fmt.Println("Value is:", x)
}
