package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	i := 3

	numbers = append(numbers[:i], numbers[i+1:]...)

	fmt.Println(numbers)
}
