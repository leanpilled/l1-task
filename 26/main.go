package main

import (
	"fmt"
	"unicode"
)

func checkUnique(s string) bool {
	set := map[rune]struct{}{}
	runes := []rune(s)

	for _, r := range runes {
		r := unicode.ToUpper(r)
		if _, ok := set[r]; ok {
			return false
		} else {
			set[r] = struct{}{}
		}
	}
	return true
}

func main() {
	fmt.Println(checkUnique("aASdf"))
}
