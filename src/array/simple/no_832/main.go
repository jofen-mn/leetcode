package main

import "fmt"

/**
翻转图像

给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。

水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。

反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。

示例 1:

输入: [[1,1,0],[1,0,1],[0,0,0]]
输出: [[1,0,0],[0,1,0],[1,1,1]]
解释: 首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
     然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
示例 2:

输入: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
输出: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
解释: 首先翻转每一行: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
     然后反转图片: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
说明:

1 <= A.length = A[0].length <= 20
0 <= A[i][j] <= 1
 */
func flipAndInvertImage(A [][]int) [][]int {
	for index := 0; index < len(A); index ++ {
		for l, r := 0, len(A[index]) - 1; l <= r; {
			temp := A[index][l]
			A[index][l] = A[index][r]
			A[index][r] = temp

			l++
			r--
		}

		for i := 0; i < len(A[index]); i++ {
			if A[index][i] == 0 {
				A[index][i] = 1
			} else if A[index][i] == 1 {
				A[index][i] = 0
			}
		}
	}

	return A
}


func main()  {
	fmt.Println("flipAndInvertImage")
	f := [][]int{{1,1,0}, {1,0,1}, {0,0,0}}
	s := [][]int{{1,1,0,0}, {1,0,0,1}, {0,1,1,1}, {1,0,1,0}}

	fmt.Println(flipAndInvertImage(f))
	fmt.Println(flipAndInvertImage(s))
}
