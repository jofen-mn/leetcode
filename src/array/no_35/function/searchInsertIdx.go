package fct

func SearchInsertIdx(nums []int, target int) int {
	var right int = len(nums) - 1
	var left int = 0
	var mid int = (left + right) / 2
	for ; right >= left; {
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return left
}
