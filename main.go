package main

import (
	"fmt"
)

func print(s string) {
	fmt.Println(s + "\n")
}

func main() {
	var s string = "Go compiles and runs! Yay!"
	print(s)
}

