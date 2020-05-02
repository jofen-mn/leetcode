package main

import (
	"fmt"
	"leetcode/src/array/simple/no_679/function"
)

func main() {
	//nums := []int{1, 2, 2, 3, 1}
	nums := []int{1,2,2,3,1,4,2}

	fmt.Println(function.FindShortestSubArray(nums))
}
