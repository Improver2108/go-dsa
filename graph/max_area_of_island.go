package graph

import (
	"fmt"

	"github.com/emirpasic/gods/v2/queues/arrayqueue"
)

// You are given a matrix grid where grid[i] is either a 0 (representing water) or 1 (representing land).

// An island is defined as a group of 1's connected horizontally or vertically. You may assume all four edges of the grid are surrounded by water.

// The area of an island is defined as the number of cells within the island.

// Return the maximum area of an island in grid. If no island exists, return 0.

// Example 1:

// Input: grid = [
//   [0,1,1,0,1],
//   [1,0,1,0,1],
//   [0,1,1,0,1],
//   [0,1,0,0,1]
// ]

// Output: 6
// Explanation: 1's cannot be connected diagonally, so the maximum area of the island is 6.

// Constraints:

// 1 <= grid.length, grid[i].length <= 50

func maxAreaOfIslandDfs(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxIslandArea := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				maxIslandArea = max(dfs(i, j), maxIslandArea)
			}
		}
	}
	return maxIslandArea
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxIslandArea := 0
	bfs := func(i, j int) int {
		queue := arrayqueue.New[[2]int]()
		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		queue.Enqueue([2]int{i, j})
		grid[i][j] = 0
		count := 0
		for !queue.Empty() {
			index, _ := queue.Dequeue()
			r, c := index[0], index[1]
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] == 0 {
					continue
				}
				queue.Enqueue([2]int{nr, nc})
				grid[nr][nc] = 0
			}
			count++
		}
		return count
	}
	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				continue
			}
			fmt.Println("maxIslandArea->", maxIslandArea)
			maxIslandArea = max(maxIslandArea, bfs(i, j))
		}
	}
	return maxIslandArea
}
func RunMaxAreaOfIsland() int {
	grid := [][]int{
		{0, 1, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 0, 1},
		{0, 1, 0, 0, 1},
	}
	return maxAreaOfIsland(grid)
}
