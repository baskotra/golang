package main

import (
	"fmt"
	"time"
)

func server1(o1 chan string) {
	time.Sleep(3 * time.Second)
	o1 <- "Server1 Started..."
}

func server2(o2 chan string) {
	time.Sleep(3 * time.Second)
	o2 <- "Server2 started..."
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
