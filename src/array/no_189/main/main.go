package main

import (
	"fmt"
	"leetcode/src/array/no_189/function"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	function.Rotate(nums, 2)

	fmt.Printf("%+v\n", nums)

	function.RotateBetter(nums, 11)

	fmt.Printf("%+v\n", nums)
}
