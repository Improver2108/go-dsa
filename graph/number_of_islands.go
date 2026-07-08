package graph

import "github.com/emirpasic/gods/v2/queues/arrayqueue"

// Given a 2D grid grid where '1' represents land and '0' represents water, count and return the number of islands.

// An island is formed by connecting adjacent lands horizontally or vertically and is surrounded by water. You may assume water is surrounding the grid (i.e., all the edges are water).

// Example 1:

// Input: grid = [
//     ["0","1","1","1","0"],
//     ["0","1","0","1","0"],
//     ["1","1","0","0","0"],
//     ["0","0","0","0","0"]
//   ]
// Output: 1
// Example 2:

// Input: grid = [
//     ["1","1","0","0","1"],
//     ["1","1","0","0","1"],
//     ["0","0","1","0","0"],
//     ["0","0","0","1","1"]
//   ]
// Output: 4
// Constraints:

// 1 <= grid.length, grid[i].length <= 100
// grid[i][j] is '0' or '1'.

func numIslandsDfs(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0
	bfs := func(i, j int) {
		queue := arrayqueue.New[[2]int]()
		queue.Enqueue([2]int{i, j})
		for !queue.Empty() {
			index, _ := queue.Dequeue()
			r, c := index[0], index[1]
			if grid[r][c] == '1' {
				if r-1 >= 0 {
					queue.Enqueue([2]int{r - 1, c})
				}
				if r+1 < m {
					queue.Enqueue([2]int{r + 1, c})
				}
				if c-1 >= 0 {
					queue.Enqueue([2]int{r, c - 1})
				}
				if c+1 < n {
					queue.Enqueue([2]int{r, c + 1})
				}
				grid[r][c] = '0'
			}
		}
	}
	for i := range m {
		for j := range n {
			if grid[i][j] == '0' {
				continue
			}
			bfs(i, j)
			count++
		}
	}
	return count
}

func RunNumIslands() int {
	grid := [][]byte{
		{'0', '1', '1', '1', '0'},
		{'0', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	return numIslands(grid)
}
