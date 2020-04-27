package function

func Generate(numRows int) [][]int {
	result := [][]int{}

	for i := 0; i < numRows; i++ {
		if i == 0 {
			result = append(result, []int{1})
		} else {
			elem := []int{}
			for k := 0; k <= i; k++ {
				if k == 0 || k == i {
					elem = append(elem, 1)
				} else {
					elem = append(elem, result[i-1][k-1]+result[i-1][k])
				}
			}
			result = append(result, elem)
		}
	}
	return result
}
