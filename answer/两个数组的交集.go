package answer

func intersection(nums1 []int, nums2 []int) []int {
	table1 := make(map[int]int)
	table2 := make(map[int]int)
	for _, num := range nums1 {
		table1[num]++
	}
	for _, num := range nums2 {
		if table1[num] > 0 {
			table2[num]++
		}
	}
	nums1 = []int{}
	for i, _ := range table2 {
		nums1 = append(nums1, i)
	}
	return nums1
}
