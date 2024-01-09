package main

import "fmt"

func quickSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)

	return arr
}

func sort(arr []int, l, r int) {
	if l < r {
		m := partition(arr, l, r)
		sort(arr, l, m-1)
		sort(arr, m+1, r)
	}
}

func partition(arr []int, l, r int) int {
	i := l - 1

	for j := l; j < r; j++ {
		if arr[j] <= arr[r] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[r] = arr[r], arr[i+1]

	return i + 1
}

func main() {
	arr := []int{1, 2, 4, 3, 5}
	fmt.Println(quickSort(arr))
}
