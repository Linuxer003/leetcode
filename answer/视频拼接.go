package answer

import "math"

func videoStitching(clips [][]int, t int) int {
	const inf = math.MaxInt64 - 1
	dp := make([]int, t+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= t; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			if l < i && i < r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[t] == inf {
		return -1
	}
	return dp[t]
}
