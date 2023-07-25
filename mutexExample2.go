package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *container) inc(name string, increment int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name] += increment
}

func main() {
	c := container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var w sync.WaitGroup

	doIncrement := func(name string, n int) {
		c.inc(name, n)
		w.Done()
	}
	w.Add(5)

	go doIncrement("a", 1)
	go doIncrement("b", 1)
	go doIncrement("b", 1)
	go doIncrement("a", 1)
	go doIncrement("b", 1)

	fmt.Println(c.counters)
}
