package main

import (
	"fmt"
)

func main() {
	dataCh := make(chan int)
	sqDataCh := make(chan int)

	arr := [5]int{1, 2, 3, 4, 5}

	go func() {
		defer close(dataCh)
		for _, v := range arr {
			dataCh <- v
		}
	}()

	go func() {
		defer close(sqDataCh)
		for n := range dataCh {
			sqDataCh <- n * n
		}
	}()

	for n := range sqDataCh {
		fmt.Println(n)
	}

}
