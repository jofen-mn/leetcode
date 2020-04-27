package main

import (
	"fmt"
	"leetcode/src/bishi/zhaoyin_1/function"
)

func main() {
	str := "*dv*fd*yyu"

	//str = function.MoveXing(str)

	//str = function.MoveXing2(str)

	str = function.MoveXing3(str)

	fmt.Println(str)

}
