package answer

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequent := make([]int, 0, upper+1)
	for _, v := range arr1 {
		frequent[v]++
	}
	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		t := frequent[v]
		frequent[v] = 0
		for i := 0; i < t; i++ {
			ans = append(ans, v)
		}
	}
	for i, v := range frequent {
		if v != 0 {
			for j := 0; j < v; j++ {
				ans = append(ans, i)
			}
		}
	}
	return ans
}

/**
如果 xx 和 yy 都出现在哈希表中，那么比较它们对应的值 \textit{rank}[x]rank[x] 和 \textit{rank}[y]rank[y]；

如果 xx 和 yy 都没有出现在哈希表中，那么比较它们本身；

对于剩余的情况，出现在哈希表中的那个元素较小。
*/
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for i, v := range arr2 {
		rank[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})
	return arr1
}
