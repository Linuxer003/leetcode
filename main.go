package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)

	if k >= n/2 {
		dp0, dp1 := 0, -prices[0]
		for i := 0; i < n; i++ {
			tmp := dp0
			if dp0 < dp1+prices[i] {
				dp0 = dp1 + prices[i]
			}
			if dp1 < tmp-prices[i] {
				dp1 = tmp - prices[i]
			}
		}
		if dp0 < dp1 {
			return dp1
		} else {
			return dp0
		}
	}

	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			if dp[j-1][1] < dp[j-1][0]-prices[i] {
				dp[j-1][1] = dp[j-1][0] - prices[i]
			}
			if dp[j][0] < dp[j-1][1]+prices[i] {
				dp[j][0] = dp[j-1][1] + prices[i]
			}
		}
	}
	return dp[k][0]
}
