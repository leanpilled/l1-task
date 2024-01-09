package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	*sync.WaitGroup
	sync.Mutex
	value int
}

func NewCounter(wg *sync.WaitGroup) *Counter {
	return &Counter{
		WaitGroup: wg,
	}
}

func (c *Counter) Increment() {
	defer c.Unlock()
	defer c.Done()
	c.Lock()
	c.value++
}

func main() {
	wg := sync.WaitGroup{}
	cnt := NewCounter(&wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go cnt.Increment()
	}

	wg.Wait()

	fmt.Println(cnt.value)

}
