package main

import (
	"fmt"
	"leetcode/src/array/simple/no_766/function"
)

func main() {
	//matrix := [][]int{{1,2,3,4}, {5,1,2,3}, {9,5,1,2}}
	matrix := [][]int{{1,2}, {2,2}}

	fmt.Println(function.IsToeplitzMatrix(matrix))
}
