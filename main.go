package main

import (
	"./logic"
	"fmt"
)

func main() {
	ints := make([]int, 100)
	iota := 0
	for i := range ints {
		iota = iota + 1
		ints[i] = iota
	}
	fmt.Println(ints)

	logic.Reverse(ints)
}
