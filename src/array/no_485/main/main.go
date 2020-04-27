package main

import (
	"fmt"
	"leetcode/src/array/no_485/function"
)

func main() {
	//nums := []int{1,0,1,1,0,1}
	nums := []int{1, 1, 0, 1, 1, 1}

	fmt.Println(function.FindMaxConsecutiveOnes(nums))
}
