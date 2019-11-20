package main

import "fmt"

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。

示例:
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
 */
func subsets(nums []int) [][]int {
	result := recursive(nums)
	return result
}

/**
通过回溯算法实现
 */
func recursive(nums []int) [][]int {
	var temp []int
	var result [][]int
	dfs(result, temp, nums, 0)
	return result
}

func dfs(result [][]int, temp []int, nums []int, index int) {
	if index == len(nums) {
		result = append(result, temp)
		return
	}
	dfs(result, temp, nums, index + 1)
	temp = append(temp, nums[index])
	dfs(result, temp, nums, index + 1)
}

func main() {
	var nums = []int{1,2,3}
	//fmt.Println(subsets(nums))
	//var nums []int
	test(nums)
	fmt.Println(nums)
}

func test(array []int) []int {
	array = append(array, 4)
	return nil
}
