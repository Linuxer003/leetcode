package answer

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
	})
	n := len(people)
	ans := make([][]int, n)
	for _, person := range people {
		spaces := person[1] + 1
		for i := 0; i < n; i++ {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}
