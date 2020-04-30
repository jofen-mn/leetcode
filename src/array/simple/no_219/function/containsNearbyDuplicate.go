package function

import (
	"math"
)

/**
 * map实现
 */
func ContainsNearbyDuplicate(nums []int, k int) bool {
	if k < 0 {
		return false
	}
	numMap := make(map[int]int, 0)
	for index, item := range nums {
		if _, ok := numMap[item]; ok && int(math.Abs(float64(numMap[item]-index))) <= k {
			return true
		} else {
			numMap[item] = index
		}
	}
	return false
}
