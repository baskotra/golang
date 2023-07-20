package main

import (
	"fmt"
	"sync"
)

func printEvenNumbers(receiveChannel chan int, w *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			receiveChannel <- i
		}
	}
	close(receiveChannel)
	w.Done()
}

func printOddNumbers(oddChannel chan int, w *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			oddChannel <- i
		}
	}
	close(oddChannel)
	w.Done()
}

func printSerial(evenChannel, oddChannel chan int, w *sync.WaitGroup) {
	fmt.Println(<-evenChannel)
	fmt.Println(<-oddChannel)
	w.Done()

}

func main() {
	var w sync.WaitGroup
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	w.Add(3)
	go printEvenNumbers(evenChannel, &w)
	go printOddNumbers(oddChannel, &w)
	go printSerial(evenChannel, oddChannel, &w)
	w.Wait()
}
