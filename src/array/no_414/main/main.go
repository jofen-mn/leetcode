package main

import (
	"fmt"
	"leetcode/src/array/no_414/function"
)

func main() {
	//nums := []int{2, 2, 3, 1}
	//nums := []int{2, 1}
	//nums := []int{2, 1, 2}
	nums := []int{1, 2, -2147483648}

	fmt.Println(function.ThirdMax(nums))
}
