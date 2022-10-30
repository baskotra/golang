package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote ", i, " to channel.")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println(v)
		time.Sleep(2 * time.Second)
	}

}
