package function


/**
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
输入: 3
输出: [1,3,3,1]
 */
func GetRow(rowIndex int) []int {
	if rowIndex > 33 {
		rowIndex = 33
	}
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	last := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		newLast := []int{1}
		for l, r := 0, 1; r < len(last);{
			newLast = append(newLast, last[l] + last[r])
			l++
			r++
		}
		newLast = append(newLast, 1)
		last = newLast
	}

	return last
}

/**
O(k)空间复杂度实现
 */
func GetRow2(rowIndex int) []int {
	if rowIndex > 33 {
		rowIndex = 33
	}
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	last := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		t := last[0]
		for r := 1; r < len(last); {
			tt := last[r]
			last[r] = t + tt
			t = tt
			r++
		}
		last = append(last, 1)
	}

	return last
}