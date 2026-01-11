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
var read chan struct{}

func main() {
	go startPolling()
	time.Sleep(15 * time.Second)
	pollEnabled = false
	go stopPoll(pollEnabled)
}

func startPolling() {
	for {
		select {
		case <-read:
			fmt.Println("Poller stopped")
		default:
			time.Sleep(5 * time.Second)
			fmt.Println(getGreeting())
		}

	}
}

func getGreeting() string {
	counter++
	return fmt.Sprintf("Hello World %d", counter)
}

func stopPoll(pollEnabled bool) {
	if pollEnabled == false {
		close(read)
	}

}
