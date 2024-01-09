package main

import "fmt"

type Set map[float64]struct{}

func main() {

	arr := [9]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -13.2}
	m := map[int]Set{}

	for _, value := range arr {
		d := getDec(value)
		s, ok := m[d]
		if ok {
			s[value] = struct{}{}
		} else {
			s = make(Set)
			s[value] = struct{}{}
			m[d] = s
		}
	}

	fmt.Println(m)

}

func getDec(num float64) int {
	return (int(num / 10)) * 10
}
