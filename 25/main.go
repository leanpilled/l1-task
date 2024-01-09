package main

import "time"

func timeSleep(tm time.Duration) {
	timer := time.NewTimer(tm)
	_ = <-timer.C
}

func main() {
	timeSleep(time.Minute)
	println("1 minute passed")
}
