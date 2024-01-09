package main

import (
	"fmt"
	"reflect"
)

func main() {
	var varBool bool
	fmt.Println(getType(varBool))
	var varInt int
	fmt.Println(getType(varInt))
	var varStr string
	fmt.Println(getType(varStr))
	varCh := make(chan int)
	fmt.Println(getType(varCh))
}

func getType(obj interface{}) reflect.Type {
	return reflect.ValueOf(obj).Type()
}
