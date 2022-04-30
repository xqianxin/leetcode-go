package search

func numEnclaves(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		helper(grid, i, 0)
	}
	for j := 0; j < len(grid[0]); j++ {
		helper(grid, 0, j)
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func helper(grid [][]int, i, j int) {
	if grid[i][j] == 0 || grid[i][j] == 2 {
		travelNext(grid, i, j)
		return
	}
	if i-1 < 0 || j-1 < 0 || i+1 == len(grid) || j+1 == len(grid[0]) {
		grid[i][j] = 2
		travelNext(grid, i, j)
		return
	}
	if grid[i-1][j] == 2 || grid[i][j-1] == 2 || grid[i+1][j] == 2 || grid[i][j+1] == 2 {
		grid[i][j] = 2
		travelNext(grid, i, j)
		return
	}
}

func travelNext(grid [][]int, i, j int) {
	if i > 0 && grid[i-1][j] == 1 {
		travelNext(grid, i-1, j)
	}
	if i < len(grid)-1 && grid[i+1][j] == 1 {
		travelNext(grid, i+1, j)
	}
	if j > 0 && grid[i][j-1] == 1 {
		travelNext(grid, i, j-1)
	}
	if j < len(grid[0]) && grid[i][j+1] == 1 {
		travelNext(grid, i, j+1)
	}
}
