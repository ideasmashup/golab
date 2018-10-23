package main

import (
	"fmt"
	"reflect"
)

func print(s string) {
	fmt.Println(s + "\n")
}

func main() {
	var s string = "Go compiles and runs! Yay!"
	var arr []string = ["a", "b", "c", "d"]

	print(s)
	print(reflect.TypeOf(arr))
}

