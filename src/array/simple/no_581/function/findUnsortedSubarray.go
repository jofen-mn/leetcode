package function

import "math"

/**
给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

你找到的子数组应是最短的，请输出它的长度。

示例 1:

输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。


说明 :

输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。
 */

/**
 * 从左到右找出最后一个破坏递增的数
 * 从右到左找出最后一个破坏递减的数
 */
func FindUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	high := 0
	low := len(nums) - 1
	max := math.MinInt32
	min := math.MaxInt32

	for index := 0; index < len(nums); index++ {
		if max <= nums[index] {
			max = nums[index]
		} else {
			high = index
		}

		if min >= nums[len(nums) - index - 1] {
			min = nums[len(nums) - index -1]
		} else {
			low = len(nums) - index - 1
		}
	}

	if high > low {
		return high - low + 1
	}

	return 0
}
