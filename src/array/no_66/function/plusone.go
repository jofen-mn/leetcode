package function

func PlusOne(dist []int) []int {
	//记录进位次数
	carryCount := 0
	if len(dist) == 0 {
		return dist
	}
	if (dist[len(dist)-1]+1)%10 == 0 {
		carryCount++
		dist[len(dist)-1] = 0
	} else {
		dist[len(dist)-1] += 1
		return dist
	}
	for i := len(dist) - 2; i >= 0; i-- {
		// 判断前一位是否发生进位
		if carryCount+i == len(dist)-1 {
			dist[i] += 1
			if dist[i]%10 == 0 {
				// 该位发生进位
				carryCount++
				// 把该位置0
				dist[i] = 0
			} else {
				return dist
			}
		}
	}

	// 判断是否要数组扩容
	if carryCount == len(dist) {
		dist = append(dist, 0)
		for i := len(dist) - 2; i >= 0; i-- {
			dist[i+1] = dist[i]
		}
		dist[0] = 1
		//newDist := []int{1}
		//newDist = append(newDist, dist...)
		return dist
	}

	return dist
}
