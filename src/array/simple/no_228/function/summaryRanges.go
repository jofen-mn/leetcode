package function

import "fmt"

/**
给定一个无重复元素的有序整数数组 nums 。
返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b

示例 1：
输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"

示例 2：
输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

示例 3：
输入：nums = []
输出：[]

示例 4：
输入：nums = [-1]
输出：["-1"]

示例 5：
输入：nums = [0]
输出：["0"]

提示：
0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
nums 中的所有值都 互不相同
*/

/**
双指针
 */
func SummaryRanges(nums []int) []string {
	if len(nums) < 0 || len(nums) > 20 {
		return []string{}
	}

	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}
	res := make([]string, 0, 10)
	for l := 0; l < len(nums); {
		r := l
		for r+1 < len(nums) && nums[r]+1 == nums[r+1] {
			r++
		}
		if r == l {
			res = append(res, fmt.Sprintf("%d", nums[l]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r]))
		}
		l = r + 1
	}

	return res
}
