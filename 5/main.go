package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	N := 5
	dataCh := make(chan int)

	go func() {
		for {
			dataCh <- rand.Intn(100)
		}
	}()

	go func() {
		for {
			fmt.Println(<-dataCh)
		}
	}()

	time.Sleep(time.Second * time.Duration(N))
}
