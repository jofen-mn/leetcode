package function

/**
 * 双指针法
 * 时间复杂度O(n)
 * 空间复杂度O(1)
 */
func TwoSum(numbers []int, target int) []int {
	var result []int
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			result = append(result, left+1, right+1)
			return result
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return result
}
