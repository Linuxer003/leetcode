package answer

func isNumber(s string) bool {
	state := [][]int{{0, 1, 2, -1, 7},
		{-1, -1, 2, -1, 7},
		{8, -1, 2, 4, 3},
		{8, -1, 3, 4, -1},
		{-1, 5, 6, -1, -1},
		{-1, -1, 6, -1, -1},
		{8, -1, 6, -1, -1},
		{-1, -1, 3, -1, -1},
		{8, -1, -1, -1, -1}}
	ans := 0
	for _, v := range s {
		if ans == -1 {
			return false
		}
		if v == ' ' {
			ans = state[ans][0]
		} else if v == '+' || v == '-' {
			ans = state[ans][1]
		} else if '0' <= v && v <= '9' {
			ans = state[ans][2]
		} else if v == 'e' {
			ans = state[ans][3]
		} else if v == '.' {
			ans = state[ans][4]
		} else {
			return false
		}
	}
	if ans == 2 || ans == 3 || ans == 6 || ans == 8 {
		return true
	}
	return false
}
