package main

import (
	"fmt"
	"leetcode/src/array/simple/no_509/function"
)

func main() {
	fmt.Println(function.Fib(1))
	fmt.Println(function.Fib(2))
	fmt.Println(function.Fib(3))
	fmt.Println(function.Fib(4))
	fmt.Println(function.Fib(5))
	fmt.Println(function.Fib(6))
	fmt.Println(function.Fib(7))

	fmt.Println(function.FibNotRecursive(5))
	fmt.Println(function.FibNotRecursive(6))
	fmt.Println(function.FibNotRecursive(7))
}
