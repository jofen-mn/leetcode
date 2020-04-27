package main

import (
	"fmt"
	"leetcode/src/tree"
)

func main() {
	nums := []int{4, 2, 6, 1, 3, 0, 0}

	//var root *tree.TreeNode
	root := tree.BuildTree(nums, 0)

	fmt.Printf("%+v\n", root)

	tree.Bianli(root)

}
