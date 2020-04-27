package function

/**
 * 二分查找法
 */
func SearchRange(nums []int, target int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-1)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] != target {
		return result
	}
	result = append(result, left)

	right = len(nums)
	for left < right {
		mid := left + (right-1)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	result = append(result, left-1)

	return result
}
