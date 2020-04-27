package main

import (
	"fmt"
	"leetcode/src/array/no_217/function"
)

func main() {
	nums := []int{1, 2, 3, 1}

	result := function.ContainsDuplicate(nums)

	fmt.Println(result)
}
