package main

import (
	"fmt"
	function2 "leetcode/src/array/simple/no_628/function"
)

func main() {
	//nums := []int{1,2,3}
	//nums := []int{1,2,-3}
	//nums := []int{1,2,3,4}
	//nums := []int{1,2,3,-4}
	nums := []int{1, 2, -3, -4}

	fmt.Println(function2.MaximumProduct(nums))
}
