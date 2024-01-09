package main

import "fmt"

func main() {
	strings := [5]string{"cat", "cat", "dog", "cat", "tree"}
	set := map[string]struct{}{}

	for _, s := range strings {
		set[s] = struct{}{}
	}

	fmt.Println(set)

}
