package answer

func PossibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]int, N+1)
	colors := make(map[int]int)
	for _, v := range dislikes {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	colors[1] = 1
	for i := 1; i <= N; i++ {
		for _, v := range graph[i] {
			if _, ok := colors[v]; !ok {
				colors[v] = -colors[i]
			} else {
				if colors[v] == colors[i] {
					return false
				}
			}
		}
	}
	return true
}
