package main

import (
	"fmt"
	"sort"
)

/*
组合总和
输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
未实现
*/
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return [][]int{}
	}
	var result = [][]int{}
	// 排序
	sort.Ints(candidates)
	fmt.Println(candidates)

	return result
}

func recursive(result [][]int, candidates []int, tmp []int, target int, index int, idx int) {
	if target == 0 {
		result[idx] = tmp
		idx++
	}
	if target < candidates[index] {
		return
	}
	recursive(result, candidates, tmp, target-candidates[index], index-1, idx)
}

func main() {
	fmt.Println("combination Sum")
	var candidates = []int{3, 1, 2, 6, 9, 4}
	combinationSum(candidates, 9)
}
