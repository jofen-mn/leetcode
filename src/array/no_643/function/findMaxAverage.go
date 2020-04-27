package function

import "math"

/**
给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

示例 1:

输入: [1,12,-5,-6,50,3], k = 4
输出: 12.75
解释: 最大平均数 (12-5-6+50)/4 = 51/4 = 12.75


注意:

1 <= k <= n <= 30,000。
所给数据范围 [-10,000，10,000]。
*/
/**
 * 滑动窗口
 */
func FindMaxAverage(nums []int, k int) float64 {
	sum := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	result := sum
	for i := k; i < len(nums); i++ {
		sum += float64(nums[i] - nums[i-k])
		result = math.Max(result, sum)
	}

	return result / float64(k)
}
