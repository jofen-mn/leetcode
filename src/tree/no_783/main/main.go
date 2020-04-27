package main

import (
	"fmt"
	"leetcode/src/tree"
	"leetcode/src/tree/no_783/function"
)

func main() {
	nums := []int{2, 1, 49, 0, 0, 13, 50, 0, 0, 0, 0}
	//nums := []int{4,2,6,1,3,0,0}

	//var root *tree.TreeNode
	root := tree.BuildTree(nums, 0)

	tree.Bianli(root)
	fmt.Println()
	tree.Zhongxubianli(root)
	fmt.Println()

	fmt.Println(function.MinDiffInBST(root))
}
