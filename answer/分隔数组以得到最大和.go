package answer

// MaxSumAfterPartitioning
func MaxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ {
		max := -1
		for j := i - 1; j >= 0 && j >= i-k; j-- {
			if arr[j] > max {
				max = arr[j]
			}
			if dp[j]+max*(i-j) > dp[i] {
				dp[i] = dp[j] + max*(i-j)
			}
		}
	}
	return dp[len(arr)]
}
