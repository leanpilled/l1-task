package main

import "fmt"

func main() {
	set1 := map[float64]struct{}{
		2.7: {}, 3.14: {}, -0.5: {}, 0: {},
	}

	set2 := map[float64]struct{}{
		2.7: {}, 3.14159: {}, 0.81828: {},
	}

	res := map[float64]struct{}{}

	for a := range set1 {
		if _, ok := set2[a]; ok {
			res[a] = struct{}{}
		}
	}

	fmt.Println(res)
}
