package main

import (
	"fmt"
	function2 "leetcode/src/array/simple/no_485/function"
)

func main() {
	//nums := []int{1,0,1,1,0,1}
	nums := []int{1, 1, 0, 1, 1, 1}

	fmt.Println(function2.FindMaxConsecutiveOnes(nums))
}
