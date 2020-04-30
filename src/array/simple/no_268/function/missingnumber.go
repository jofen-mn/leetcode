package function

/**
 * 数学运算
 */
func MissingNumber(nums []int) int {
	result := 0
	max := 0
	min := -1
	for _, item := range nums {
		if item > max {
			max = item
		}
		if item == 0 {
			min = item
		}
		result += item
	}

	if min != 0 {
		return 0
	}

	if result == max*(max+1)/2 {
		return max + 1
	}

	return max*(max+1)/2 - result
}
