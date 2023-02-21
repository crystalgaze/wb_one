package main

import (
	"fmt"
	"sync"
)

var group sync.WaitGroup

type Counter struct {
	mutex sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
	}
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
	defer group.Done()
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	с := NewCounter()

	for i := 0; i < 100; i++ {
		group.Add(1)
		go с.Increment()
	}
	group.Wait()
	fmt.Println(с.Value())
}
