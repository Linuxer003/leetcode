package answer

func KClosest(points [][]int, K int) [][]int {
	if len(points) <= 1 {
		return points
	}
	quickSort(points)
	return points[:K]
}

func quickSort(values [][]int) {
	if len(values) <= 1 {
		return
	}
	left := values[0]
	temp := values[0][0]*values[0][0] + values[0][1]*values[0][1]
	l, r := 0, len(values)-1
	for l < r {
		for r > 0 && values[r][0]*values[r][0]+values[r][1]*values[r][1] >= temp {
			r--
		}
		if r > l {
			values[l] = values[r]
		}
		for l < r && values[l][0]*values[l][0]+values[l][1]*values[l][1] < temp {
			l++
		}
		if l < r {
			values[r] = values[l]
		}
	}
	values[l] = left
	quickSort(values[:l])
	quickSort(values[l+1:])
}
