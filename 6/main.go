package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func oneWay(quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func secWay(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Add(1)
	quit := make(chan bool)
	go oneWay(quit, &wg)
	time.Sleep(3 * time.Second)
	quit <- true
	wg.Wait()

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go secWay(ctx, &wg)
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}
