package answer

func uniqueOccurrences(arr []int) bool {
	table1 := make(map[int]int)
	table2 := make(map[int]int)

	for _, num := range arr {
		table1[num]++
	}
	for _, num := range table1 {
		table2[num]++
	}
	for _, num := range table2 {
		if num != 1 {
			return false
		}
	}
	return true
}
