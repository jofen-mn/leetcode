package function

import "math"

/**

 */
/**
 * 将数组元素数值减1作为数组下标，将该下标对应的数值变成负数
 */
func FindDisappearedNumbers(nums []int) []int {
	var result []int
	for _, item := range nums {
		nums[int(math.Abs(float64(item)))-1] = -int(math.Abs(float64(nums[int(math.Abs(float64(item)))-1])))
	}

	for index, _ := range nums {
		if nums[index] > 0 {
			result = append(result, index+1)
		}
	}

	return result
}
