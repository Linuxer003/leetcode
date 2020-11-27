package answer

func makeConnected(n int, connections [][]int) int {
	edges := make([][]int, n)
	used := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		used[u] = true
		for _, v := range edges[u] {
			if !used[v] {
				dfs(v)
			}
		}
	}

	if len(connections) < n-1 {
		return -1
	}
	for _, v := range connections {
		edges[v[0]] = append(edges[v[0]], v[1])
		edges[v[1]] = append(edges[v[1]], v[0])
	}

	part := 0
	for i := 0; i < n; i++ {
		if !used[i] {
			part++
			dfs(i)
		}
	}
	return part - 1
}
