package main

import (
	"fmt"
	"leetcode/src/array/no_66/function"
)

func main() {
	dist := []int{9, 9, 9, 9}
	dist = function.PlusOne(dist)

	fmt.Println(len(dist))
	fmt.Printf("%+v\n", dist)
}
