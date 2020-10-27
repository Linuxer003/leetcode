package answer

func closedIsland(grid [][]int) int {
	num, rowNum, colNum := 0, len(grid), len(grid[0])
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for row, v1 := range grid {
		for col, v2 := range v1 {
			if v2 == 1 {
				continue
			}
			closed := true
			queue := [][]int{{row, col}}
			if row == 0 || row == rowNum-1 || col == 0 || col == colNum-1 {
				closed = false
			}
			for len(queue) != 0 {
				head := queue[0]
				queue = queue[1:]
				for _, v := range dir {
					newRow, newCol := head[0]+v[0], head[1]+v[1]
					if newRow < 0 || newRow >= rowNum || newCol < 0 || newCol >= colNum {
						continue
					}
					if grid[newRow][newCol] == 1 {
						continue
					}
					if newRow == 0 || newRow == rowNum-1 || newCol == 0 || newCol == colNum-1 {
						closed = false
					}
					queue = append(queue, []int{newRow, newCol})
					grid[newRow][newCol] = 1
				}
			}
			if closed {
				num++
			}
		}
	}
	return num
}
