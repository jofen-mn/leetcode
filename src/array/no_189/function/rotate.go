package function

/**
 * 暴力解决
 * 时间复杂度O(kn)
 * 空间复杂度O(1)
 */
func Rotate(nums []int, k int) {
	if k <= 0 {
		return
	}

	k %= len(nums)

	var temp int
	for i := 0; i < k; i++ {
		temp = nums[len(nums)-1]
		for index := len(nums) - 1; index > 0; index-- {
			nums[index] = nums[index-1]
		}

		nums[0] = temp
	}
}

/**
 *
 * 时间复杂度O(n)
 * 空间复杂度O(1)
 */
func RotateBetter(nums []int, k int) {
	if len(nums) < 2 {
		return
	}

	k %= len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		var temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp

		start++
		end--
	}
}
