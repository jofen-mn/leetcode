package main

import (
	"fmt"
	function2 "leetcode/src/array/simple/no_219/function"
)

func main() {
	nums := []int{1, 0, 1, 1}

	result := function2.ContainsNearbyDuplicate(nums, 1)

	fmt.Println(result)
}
