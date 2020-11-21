package answer

import "math/rand"

func rand10() int {
	var row, col, idx int
	for true {
		row = rand7()
		col = rand7()
		idx = col + (row-1)*7
		if idx <= 40 {
			break
		}
	}
	return 1 + (idx-1)%10
}

func rand7() int {
	return 1 + rand.Intn(7)
}
