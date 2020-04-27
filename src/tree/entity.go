package tree

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 默认0为null
 */
func BuildTree(nums []int, index int) *TreeNode {
	if len(nums) == 0 || index >= len(nums) || nums[index] == 0 {
		return nil
	}

	tree := &TreeNode{nums[index], nil, nil}
	left := BuildTree(nums, 2*(index+1)-1)
	right := BuildTree(nums, 2*(index+1))

	tree.Left = left
	tree.Right = right
	return tree
}

func Bianli(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Printf("%d ", tree.Val)
	Bianli(tree.Left)
	Bianli(tree.Right)
}

func Zhongxubianli(tree *TreeNode) {
	if tree == nil {
		return
	}

	Zhongxubianli(tree.Left)
	fmt.Printf("%d ", tree.Val)
	Zhongxubianli(tree.Right)
}
