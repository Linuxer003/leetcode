package answer

func fourSumCount(A, B, C, D []int) int {
	sum := map[int]int{}

	for _, a := range A {
		for _, b := range B {
			sum[a+b]++
		}
	}

	count := 0
	for _, c := range C {
		for _, d := range D {
			count += sum[-(c + d)]
		}
	}

	return count
}
