package function

/**
给定一个未经排序的整数数组，找到最长且连续的的递增序列。

示例 1:

输入: [1,3,5,4,7]
输出: 3
解释: 最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
示例 2:

输入: [2,2,2,2,2]
输出: 1
解释: 最长连续递增序列是 [2], 长度为1。
 */
/**
双指针
 */
func FindLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := 0
	result := 0
	for ; right < len(nums) - 1; {
		if nums[right] >= nums[right + 1] {
			lenght := right - left + 1
			if result < lenght {
				result = lenght
			}
			left = right + 1
		}

		right++
	}

	if lenght := right - left + 1; lenght > result {
		return lenght
	}

	return result
}
