package main

import (
	"fmt"
	"leetcode/src/array/no_219/function"
)

func main() {
	nums := []int{1, 0, 1, 1}

	result := function.ContainsNearbyDuplicate(nums, 1)

	fmt.Println(result)
}
