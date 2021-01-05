package main

import (
	"fmt"
	"leetcode/src/array/simple/no_746/function"
)

func main() {
	cost := []int{10, 15, 20}
	//cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	fmt.Println(function.MinCostClimbingStairs(cost))
}
