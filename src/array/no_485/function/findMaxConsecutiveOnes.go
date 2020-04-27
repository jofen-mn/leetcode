package function

/**
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
注意：

输入的数组只包含 0 和1。
输入数组的长度是正整数，且不超过 10,000。
*/
/**
 * 双指针
 */
func FindMaxConsecutiveOnes(nums []int) int {
	left := 0
	right := 0
	max := 0

	var val int
	for right < len(nums) {
		if right == len(nums)-1 {
			val = right - left + 1
		}
		if nums[right] != 1 {
			val = right - left
			left = right + 1
		}

		if val > max {
			max = val
		}

		right++
	}

	return max
}
