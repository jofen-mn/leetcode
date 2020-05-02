package main

import (
	"fmt"
	"leetcode/src/array/simple/no_724/function"
)

func main() {
	//nums := []int{1, 7, 3, 6, 5, 6}
	//nums := []int{1, 7, 6, 5, 6}
	//nums := []int{1, 2, 3}
	//nums := []int{}
	//nums := []int{-1, -1, -1, -1, -1, 0}
	//nums := []int{0,-1, -1, -1, -1, -1}
	//nums := []int{-1, -1, 0, -1, -1, -1}
	nums := []int{-1, -1, 0, -1, -1, 0}

	fmt.Println(function.PivotIndex(nums))
}
