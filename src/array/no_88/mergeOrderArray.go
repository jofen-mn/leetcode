package main

import (
	"fmt"
	"sort"
)
/**

 */

/**
暴力法实现，双重循环比较

 */
func mergeOrderArray(nums1 []int, m int, nums2 []int, n int)  {
	for i :=0; i < n; i++ {
		j := 0
		for ; j < m; {
			if nums2[i] >= nums1[j] {
				j++
			} else {
				//插入
				for k := m; k > j; k-- {
					nums1[k] = nums1[k - 1]
				}
				nums1[j] = nums2[i]
				break
			}
		}
		if j == m && m < len(nums1){
			nums1[m] = nums2[i]
		}
		m++
	}
	printArr(nums1)
}
/**
借助辅助空间，不推荐，速度也没见提升多少，还占用内存空间
 */
func merge2(nums1 []int, m int, nums2 []int, n int) {
	tmpArr := make([]int,m+n)
	i := 0
	j := 0
	k := 0
	for ; i < m && j < n; {
		if nums1[i] < nums2[j] {
			tmpArr[k] = nums1[i]
			i++
		} else {
			tmpArr[k] = nums2[j]
			j++
		}
		k++
	}
	for ; i < m; i++ {
		tmpArr[k] = nums1[i]
		k++
	}
	for ; j < n; j++ {
		tmpArr[k] = nums2[j]
		k++
	}
	index := 0
	for ; index < len(tmpArr); index++ {
		nums1[index] = tmpArr[index]
	}
	for ; index < len(nums1); index++ {
		nums1[index] = 0
	}
	printArr(tmpArr)
}
/**
通过go的排序api来完成，不推荐，毕竟是做算法题（虽然如此，但是速度（leetcode给出的速度）比我自己写的快^^|）
不知道是网络原因还是什么
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := m; i < m + n && i < len(nums1) && j < n; i++ {
		nums1[i] = nums2[j]
		j++
	}
	sort.Ints(nums1)
	printArr(nums1)
}
/**
用于测试
 */
func printArr(nums []int) {
	for i := 0; i < len(nums) && nums[i] != 0; i++ {
		fmt.Print(nums[i])
		fmt.Print(" ")
	}
}

func main () {
	nums1 := []int {1,2,3,0,0,0}
	//nums2 := []int {2,4,6,8}[1,2,3,0,0,0]
	nums2 := []int {2,5,6}
	merge(nums1, len(nums1) - len(nums2), nums2, len(nums2))
}
