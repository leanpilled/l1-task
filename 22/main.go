package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigY := big.NewFloat(2e25)
	bigX := big.NewFloat(2e25)
	res := new(big.Float)

	fmt.Println(res.Quo(bigY, bigX))
	fmt.Println(res.Add(bigY, bigX))
	fmt.Println(res.Sub(bigY, bigX))
	fmt.Println(res.Mul(bigY, bigX))
}
