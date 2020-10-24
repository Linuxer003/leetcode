package answer

import (
	"sort"
)

func kthSmallest(mat [][]int, k int) int {
	var ans []int
	ans = mat[0]

	for i := 1; i < len(mat); i++ {
		var st []int
		for j := 0; j < len(mat[i]); j++ {
			for t := 0; t < len(ans); t++ {
				st = append(st, ans[t]+mat[i][j])
			}
		}
		sort.Ints(st)
		ans = st
		if len(ans) > k {
			ans = ans[:k]
		} else {
			ans = ans[:]
		}
	}

	return ans[k-1]
}
