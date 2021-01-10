package main

import "fmt"

/**
翻转矩阵

给定一个矩阵 A， 返回 A 的转置矩阵。
矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

 
示例 1：
输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]
示例 2：

输入：[[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]
 

提示：

1 <= A.length <= 1000
1 <= A[0].length <= 1000
 */
func transpose(A [][]int) [][]int {
	var res [][]int
	for i := 0; i < len(A[0]); i++ {
		t := make([]int, 0, len(A))
		for j := 0; j < len(A); j++ {
			t = append(t, A[j][i])
		}
		res = append(res, t)
	}

	return res
}

func main()  {
	f := [][]int{{1,2,3}, {4,5,6}, {7, 8, 9}}
	s := [][]int{{1,2,3}, {4,5,6}}

	fmt.Println(transpose(f))
	fmt.Println(transpose(s))
}