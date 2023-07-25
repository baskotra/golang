package main

import "fmt"

func sendOnly(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	go sendOnly(ch)
	for i := 0; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}
