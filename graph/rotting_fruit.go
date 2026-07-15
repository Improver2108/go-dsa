package graph

// You are given a 2-D matrix grid. Each cell can have one of three possible values:

// 0 representing an empty cell
// 1 representing a fresh fruit
// 2 representing a rotten fruit
// Every minute, if a fresh fruit is horizontally or vertically adjacent to a rotten fruit, then the fresh fruit also becomes rotten.

// Return the minimum number of minutes that must elapse until there are zero fresh fruits remaining. If this state is impossible within the grid, return -1.

// Example 1:
// Input: grid = [[1,1,0],[0,1,1],[0,1,2]]
// Output: 4

// Example 2:
// Input: grid = [[1,0,1],[0,2,0],[1,0,1]]
// Output: -1

// Constraints:

// 1 <= grid.length, grid[i].length <= 10

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][3]int{}
	totalFresh := 0
	for i := range m {
		for j := range n {
			switch grid[i][j] {
			case 1:
				totalFresh++
			case 2:
				queue = append(queue, [3]int{i, j, 0})
			}
		}
	}

	if totalFresh == 0 {
		return 0
	}

	directions := [...][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dist := 0
	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		dist = queue[0][2]
		queue = queue[1:]
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] != 1 {
				continue
			}
			grid[nr][nc] = 2
			totalFresh--
			queue = append(queue, [3]int{nr, nc, dist + 1})
		}
	}
	if totalFresh == 0 {
		return dist
	}
	return -1
}

func RunOrangeRotting() int {
	grid := [][]int{{1, 1, 0}, {0, 1, 1}, {0, 1, 2}}
	return orangesRotting(grid)
}
