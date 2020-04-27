package main

import (
	"fmt"
	"leetcode/src/array/no_167/function"
)

func main() {
	numbers := []int{2, 7, 11, 15}

	result := function.TwoSum(numbers, 100)

	fmt.Printf("%+v\n", result)
}
