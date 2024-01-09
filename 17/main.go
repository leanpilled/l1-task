package main

import "fmt"

func bSearch(arr []int, x int) (int, bool) {
	l := 0
	r := len(arr) - 1
	var m int

	for l <= r {
		m = (l + r) / 2

		if arr[m] < x {
			l = m + 1
		} else if arr[m] > x {
			r = m - 1
		} else {
			return m, true
		}
	}

	return 0, false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 19}

	if idx, ok := bSearch(arr, 19); ok {
		fmt.Println(idx)
	} else {
		fmt.Println("Element not found")
	}
}
