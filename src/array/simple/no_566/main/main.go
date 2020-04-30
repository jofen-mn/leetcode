package main

import (
	"fmt"
	"leetcode/src/array/simple/no_566/function"
)

func main() {
	nums := [][]int{{1,2}, {3,4}}


	//fmt.Println(function.MatrixReshape(nums, 1, 4))
	//fmt.Println(function.MatrixReshape(nums, 1, 3))
	fmt.Println(function.MatrixReshape(nums, 2, 4))
}
