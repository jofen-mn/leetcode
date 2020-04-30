package main

import (
	"fmt"
	"leetcode/src/array/simple/no_661/function"
)

func main() {
	M := [][]int{{1,1,1}, {1,0,1}, {1,1,1}}

	fmt.Println(function.ImageSmoother(M))
}
