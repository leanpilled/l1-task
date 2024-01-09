package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func worker(ch chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {

	dataCh := make(chan int)

	N := 5

	for i := 0; i < N; i++ {
		go worker(dataCh)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			close(dataCh)
			fmt.Println("exiting")
			return
		default:
			dataCh <- rand.Intn(100)
		}
	}
}
