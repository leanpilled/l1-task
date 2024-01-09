package main

import (
	"strings"
)

var justString string

func createHugeString(num int) string {
	char := "А"
	str := strings.Repeat(char, num)
	return str
}

func correctFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	correctFunc()
}
