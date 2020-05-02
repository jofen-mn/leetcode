package function

/**
给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中所有的 i (1 <= i < n)，总满足 array[i] <= array[i + 1]。

 

示例 1:

输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
示例 2:

输入: nums = [4,2,1]
输出: false
解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
 */
/**
遍历数组，每次比较三个数，开始时第一个数已经设置为大于前一个数，因此需要从index等于1开始，
若中间的数比第三个数大，将计数器加1,判断计数器值，然后比较第一个数是否比第三个数大，若是则将第三个数变大
若不是将中间的数变小
 */
func CheckPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}

	count := 0
	if nums[0] > nums[1] {
		nums[0] = nums[1]
		count++
	}

	for index := 1; index < len(nums) - 1; index++ {
		right := nums[index+1]
		if nums[index] > right {
			count++
			if count > 1 {
				return false
			}

			left := nums[index - 1]
			if left > right {
				nums[index + 1] = nums[index]
			} else {
				nums[index] = left
			}
		}
	}

	return count < 2
}
