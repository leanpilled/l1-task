package main

import (
	"fmt"
	"sync"
)

func square(x int) int {
	return x * x
}

func main() {
	dataCh := make(chan int)

	arr := [5]int{1, 2, 3, 4, 5}

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < len(arr); i++ {
			wg.Add(1)
			go func(x int) {
				defer wg.Done()
				res := square(x)
				dataCh <- res
			}(arr[i])
		}
		wg.Wait()
		close(dataCh)
	}()

	var sum int
	for n := range dataCh {
		sum += n
	}
	fmt.Println(sum)

}
