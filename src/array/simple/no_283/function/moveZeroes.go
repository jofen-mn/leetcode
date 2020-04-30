package function

/**
 * 双指针
 * 时间复杂度O(n)
 */
func MoveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	left := 0
	right := left

	for right < len(nums) {
		if nums[left] != 0 {
			left++
			if left > right {
				right = left
			}
		} else {
			for right < len(nums) && nums[right] == 0 {
				right++
			}

			if right < len(nums) {
				// swap
				nums[left] = nums[right]
				nums[right] = 0
				left++
				right++
			}
		}
	}
}
