package answer

func smallerNumbersThanCurrent(nums []int) []int {
	h := make([]int, 101)

	for _, n := range nums {
		h[n]++
	}
	for i := 1; i <= 100; i++ {
		h[i] += h[i-1]
	}

	ans := make([]int, len(nums))
	for i, n := range nums {
		if n > 0 {
			ans[i] = h[n-1]
		}
	}
	return ans
}
