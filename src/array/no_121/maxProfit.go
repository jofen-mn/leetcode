package main

import (
	"fmt"
	"math"
)

/**
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 */
func maxProfit(prices []int) int {
	return byDP(prices)
}

/**
双指针实现，遍历数组，通过双指针扫描后面的数组元素，实时记录最大利润，
若最后结果为负数，则利润为0
执行用时：240ms左右
 */
func byTwoPoint(prices []int) int {
	len := len(prices)
	maxPrices := 0
	if len == 1 {
		return maxPrices
	} else if len == 2 {
		if prices[0] < prices[1] {
			return prices[1] - prices[0]
		} else {
			return maxPrices
		}
	}
	for i := 0; i < len; i++ {
		left := i + 1
		right := len - 1
		for ; left <= right;  {
			tmp := 0
			if prices[left] > prices[right] {
				tmp = prices[left] - prices[i]
			} else {
				tmp = prices[right] - prices[i]
			}
			if maxPrices < tmp {
				maxPrices = tmp
			}
			left ++
			right --
		}
	}
	if maxPrices > 0 {
		return maxPrices
	} else {
		return 0
	}
}

/**
通过动态规划实现
记录一个买入的最小值，后续对比交换该最小值，同时计算收益值是否大于当前的记录的最大收益值
当查询到一个新的买入的较小值，将该较小值更新到最小值记录，因为当前买入值是记录的最小值，即最大收益记录将不会更新
 */
func byDP(prices []int) int {
	inMin := math.MaxInt64
	proMax := 0
	for i := 0; i < len(prices); i++ {
		if inMin > prices[i] {
			inMin = prices[i]
		}
		if prices[i] - inMin > proMax {
			proMax = prices[i] - inMin
		}
	}
	return proMax
}

func main() {
	prices := []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))
}