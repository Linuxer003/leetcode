package algorithm

// radixSort radix sort algorithm implement
// this sort algorithm spent the linear time
// radix sort can fix the disadvantage of space usage in counting sort
// when the max number in array is vary large.
func radixSort(nums []int) []int {
	n := len(nums)
	buf := make([]int, n)
	maxVal := max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		// this part is use counting sort
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		// the nums is sorted on some radix now
		// after some cycles, the nums is complete sorted
		copy(nums, buf)
	}
	return nums
}

// max
// find the max number in a
func max(a ...int) int {
	ans := a[0]
	for _, v := range a[1:] {
		if v > ans {
			ans = v
		}
	}
	return ans
}
