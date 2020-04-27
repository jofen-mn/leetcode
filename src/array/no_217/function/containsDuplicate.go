package function

/**
 * 利用map
 * 时间复杂度O(n)
 * 空间复杂度O(n)
 */
func ContainsDuplicate(nums []int) bool {
	numMap := make(map[int]interface{}, 0)
	for _, item := range nums {
		if _, ok := numMap[item]; ok {
			return true
		} else {
			numMap[item] = ""
		}
	}
	return false
}
