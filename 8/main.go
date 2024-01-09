package main

import (
	"fmt"
	"strconv"
)

func changeBit(num int64, bit int, idx int) int64 {
	mask := int64(1 << idx)
	currBit := num & mask

	if (currBit == 1 && bit == 1) || (currBit == 0 && bit == 0) {
		return num
	}

	if bit == 1 {
		num |= 1 << idx
	} else if bit == 0 {
		num &= ^(1 << idx)
	}

	return num

}

func main() {
	var num int64 = 1125899906842624

	fmt.Println(strconv.FormatInt(num, 2))

	fmt.Println(changeBit(num, 0, 50))
}
