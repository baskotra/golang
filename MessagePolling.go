/*
StartMessagePoller, StopMessagePoller
GetMessage() -> return message -> Mesage 1, Message 2, Message 3 and so on
PollEnabled := true -> Call StartMessagePoller -> fetch message every 5 seconds
After 15 seconds, pollEnabled flag will be set to false
pollEnabled := false -> call StopMessagePoller -> poller should be stopped and Print poller Stopped
*/

package main

import (
	"fmt"
	"time"
)

var pollEnabled = true
var counter int32 = 0
//initialize channel, since a nil channel always blocks forever
var read = make(chan struct{})

func main() {
	go startPolling()
	time.Sleep(5 * time.Second)
	pollEnabled = false
	go stopPoll(pollEnabled)
	time.Sleep(1 * time.Second)
}

func startPolling() {
	for {
		select {
		case <-read:
			fmt.Println("Poller stopped")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(getGreeting())
		}

	}
}

func getGreeting() string {
	counter++
	return fmt.Sprintf("Hello World %d", counter)
}

func stopPoll(pollEnabled bool) {
	if !pollEnabled {
		close(read)
	}

}
