package answer

//islandPerimeter 统计岛屿周长
func islandPerimeter(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}
	var land, double int
	var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i, v1 := range grid {
		for j, v2 := range v1 {
			if v2 == 1 {
				land++
				for _, v3 := range dir {
					if i+v3[0] < 0 || i+v3[0] >= row || j+v3[1] < 0 || j+v3[1] >= col {
						continue
					}
					if grid[i+v3[0]][j+v3[1]] == 1 {
						double++
					}
				}
			}
		}
	}
	return land*4 - double
}
