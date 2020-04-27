package main

import (
	"fmt"
	"leetcode/src/array/no_283/function"
)

func main() {
	nums := []int{0, 0}

	function.MoveZeroes(nums)

	fmt.Printf("%+v\n", nums)
}
