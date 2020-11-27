package algorithm

// countSort
// this sort algorithm spend the linear time
func countSort(a []int) []int {
	n := len(a)
	k := 0
	for i := 0; i < n; i++ {
		if a[i] > k {
			k = a[i]
		}
	}
	c := make([]int, k+1)
	b := make([]int, n)
	for _, v := range a {
		c[v]++
	}
	// c[i] now contains the number of elements equal to i

	for i := 1; i <= k; i++ {
		c[i] += c[i-1]
	}
	// c[i] now contains the number of elements less than or equal to i

	for i := n - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]]--
	}
	// now b is sorted

	return b
}
