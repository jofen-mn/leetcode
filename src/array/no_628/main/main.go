package main

import (
	"fmt"
	"leetcode/src/array/no_628/function"
)

func main() {
	//nums := []int{1,2,3}
	//nums := []int{1,2,-3}
	//nums := []int{1,2,3,4}
	//nums := []int{1,2,3,-4}
	nums := []int{1, 2, -3, -4}

	fmt.Println(function.MaximumProduct(nums))
}
