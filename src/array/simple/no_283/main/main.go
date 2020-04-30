package main

import (
	"fmt"
	function2 "leetcode/src/array/simple/no_283/function"
)

func main() {
	nums := []int{0, 0}

	function2.MoveZeroes(nums)

	fmt.Printf("%+v\n", nums)
}
