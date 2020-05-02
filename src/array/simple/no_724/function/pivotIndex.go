package function

import "fmt"

/**
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

示例 1:
输入:
nums = [1, 7, 3, 6, 5, 6]
输出: 3
解释:
索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
同时, 3 也是第一个符合要求的中心索引。

示例 2:
输入:
nums = [1, 2, 3]
输出: -1
解释:
数组中不存在满足此条件的中心索引。
 */

/**
双指针（没解出来）
 */
func PivotIndexI(nums []int) int {
	left := 0
	right := len(nums) - 1
	leftSum := 0
	rightSum := 0

	for ; left < right; {
		if leftSum > rightSum {
			if nums[right] >= 0 {
				rightSum += nums[right]
				right--
			} else {
				leftSum += nums[left]
				left++
			}
			//rightSum += nums[right]
			//right--
		} else {
			if nums[left] >= 0 {
				leftSum += nums[left]
				left++
			} else {
				rightSum += nums[right]
				right--
			}
			//leftSum += nums[left]
			//left++
		}
	}

	fmt.Printf("%d\t%d\t%d\t%d\n", left, right, leftSum, rightSum)
	if left >= right && leftSum == rightSum {
		return right
	}

	return -1
}

/**
前缀和（官方提供的解法）
 */
func PivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for index := 0; index < len(nums); index++ {
		if leftSum == sum - leftSum - nums[index] {
			return index
		}

		leftSum += nums[index]
	}

	return -1
}



