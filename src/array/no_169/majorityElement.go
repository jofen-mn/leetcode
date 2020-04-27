package main

import (
	"fmt"
	"sort"
)

/**
给定一个大小为 n 的数组，找到其中的众数。
众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在众数。
示例 1:
输入: [3,2,3]
输出: 3
示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
*/

/**
通过api
*/
func majorityElementByApi(num []int) int {
	sort.Ints(num)
	mid := len(num) / 2
	return num[mid]
}

/**
自己实现算法
摩尔投票法
定义一个计数器从数组第一个位置开始计数，置为1
遇到相同的加1，不同的减1，直到数组遍历结束或者计数器为零
计数器为0交换下一个元素
*/
func majorityElement(num []int) int {
	count := 1
	result := num[0]
	for i := 1; i < len(num); i++ {
		if result == num[i] {
			count++
		} else {
			count--
			if count == 0 {
				result = num[i+1]
			}
		}
	}
	return result
}

func main() {
	var num = []int{2, 3, 4, 2, 2, 1, 6, 2, 2}
	fmt.Println(majorityElement(num))
}
