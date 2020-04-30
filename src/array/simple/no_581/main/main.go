package main

import (
	"fmt"
	"leetcode/src/array/simple/no_581/function"
)

func main() {
	nums := []int {2, 6, 4, 8, 10, 9, 15}
	//nums := []int {2, 1}
	//nums := []int {1, 3, 2, 2, 2}

	fmt.Println(function.FindUnsortedSubarray(nums))
}
