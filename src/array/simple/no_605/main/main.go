package main

import (
	"fmt"
	"leetcode/src/array/simple/no_605/function"
)

func main() {
	//nums := []int{1,0,0,0,1}
	nums := []int{0,0,1,0,0} // [0,0,1,0,0]

	fmt.Println(function.CanPlaceFlowers(nums, 1))
	fmt.Println(function.CanPlaceFlowers(nums, 2))
}
