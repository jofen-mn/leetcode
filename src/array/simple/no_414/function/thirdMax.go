package function

import (
	"fmt"
	"math"
)

/**
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。

示例 1:

输入: [3, 2, 1]

输出: 1

解释: 第三大的数是 1.
示例 2:

输入: [1, 2]

输出: 2

解释: 第三大的数不存在, 所以返回最大的数 2 .
示例 3:

输入: [2, 2, 3, 1]

输出: 1

解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。
*/
/**
 *
 */
func ThirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}

	max1 := math.MinInt32
	max2 := max1
	max3 := max1
	// 去重
	numsMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			numsMap[num]++
		} else {
			numsMap[num] = 1
		}
	}

	for num, _ := range numsMap {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}
	}

	fmt.Printf("%d\t%d\t%d\t", max1, max2, max3)
	if len(numsMap) == 2 {
		return max1
	}

	return max3
}
