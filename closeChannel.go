package main

import "fmt"

func numbers(loopChan chan int) {
	for i := 0; i <= 5; i++ {
		loopChan <- i
	}
	close(loopChan)
}

func main() {
	loopChan := make(chan int)
	go numbers(loopChan)
	for {
		value, ok := <-loopChan
		if ok == false {
			break
		}
		fmt.Println("Value is: ", value)
	}

}
