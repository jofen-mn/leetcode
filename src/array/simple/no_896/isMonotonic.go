package main

import "fmt"

/**
如果数组是单调递增或单调递减的，那么它是单调的。
如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。
如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。
当给定的数组 A 是单调数组时返回 true，否则返回 false。

 

示例 1：
输入：[1,2,2,3]
输出：true

示例 2：
输入：[6,5,4,4]
输出：true

示例 3：
输入：[1,3,2]
输出：false

示例 4：
输入：[1,2,4,5]
输出：true

示例 5：
输入：[1,1,1]
输出：true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/monotonic-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func isMonotonic(A []int) bool {
	inc, dec := false, false
	for index := 0; index < len(A) - 1; index++ {
		if A[index] < A[index+1] {
			inc = true
		}
		if A[index] > A[index+1] {
			dec = true
		}
	}

	return !(inc&&dec)
}

func main()  {
	f := []int{1,2,2,3}
	f2 := []int{6,5,4,4}
	f3 := []int{1,3,2}
	f4 := []int{1,2,4,5}

	fmt.Println(isMonotonic(f))
	fmt.Println(isMonotonic(f2))
	fmt.Println(isMonotonic(f3))
	fmt.Println(isMonotonic(f4))

}
