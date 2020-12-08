package answer

import "math"

func splitIntoFibonacci(S string) (F []int) {
	n := len(S)
	var backtrace func(index, sum, prev int) bool
	backtrace = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}
		cur := 0
		for i := index; i < n; i++ {
			if i > index && S[index] == '0' {
				break
			}
			cur = cur*10 + int(S[i]-'0')
			if cur > math.MaxInt32 {
				break
			}

			if len(F) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}
			F = append(F, cur)
			if backtrace(i+1, prev+cur, cur) {
				return true
			}
			F = F[:len(F)-1]
		}
		return false
	}
	backtrace(0, 0, 0)
	return
}
