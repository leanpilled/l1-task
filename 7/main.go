package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}
	wg := sync.WaitGroup{}
	var mutex sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		for i := 0; i < 5; i++ {
			m[i] = i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		for i := 0; i < 10; i++ {
			m[i] = -i
		}
	}()

	wg.Wait()

	fmt.Println(m)

}
