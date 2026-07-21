package graph

import (
	"maps"
	"slices"
)

// You are given a rectangular island heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

// The islands borders the Pacific Ocean from the top and left sides, and borders the Atlantic Ocean from the bottom and right sides.

// Water can flow in four directions (up, down, left, or right) from a cell to a neighboring cell with height equal or lower. Water can also flow into the ocean from cells adjacent to the ocean.

// Find all cells where water can flow from that cell to both the Pacific and Atlantic oceans. Return it as a 2D list where each element is a list [r, c] representing the row and column of the cell. You may return the answer in any order.

// Example 1:

// Input: heights = [
//   [4,2,7,3,4],
//   [7,4,6,4,7],
//   [6,3,5,3,6]
// ]

// Output: [[0,2],[0,4],[1,0],[1,1],[1,2],[1,3],[1,4],[2,0]]
// Example 2:

// Input: heights = [[1],[1]]

// Output: [[0,0],[1,0]]
// Constraints:

// 1 <= heights.length, heights[r].length <= 100
// 0 <= heights[r][c] <= 1000

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacificMap := make(map[[2]int]struct{})
	atlanticMap := make(map[[2]int]struct{})

	for i := range m {
		for j := range n {
			if i == 0 || j == 0 {
				pacificMap[[2]int{i, j}] = struct{}{}
			}
			if i == m-1 || j == n-1 {
				atlanticMap[[2]int{i, j}] = struct{}{}
			}
		}
	}
	bfs := func(mp map[[2]int]struct{}) {
		queue := slices.Collect(maps.Keys(mp))
		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for len(queue) > 0 {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				_, exists := mp[[2]int{nr, nc}]
				if nr < 0 || nr >= m || nc < 0 || nc >= n || heights[r][c] > heights[nr][nc] || exists {
					continue
				}
				mp[[2]int{nr, nc}] = struct{}{}
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
	bfs(pacificMap)
	bfs(atlanticMap)

	res := [][]int{}

	for key, _ := range pacificMap {
		if _, exists := atlanticMap[key]; exists {
			res = append(res, key[:])
		}
	}

	return res
}

func RunPacificAtlantic() [][]int {
	heights := [][]int{{4, 2, 7, 3, 4}, {7, 4, 6, 4, 7}, {6, 3, 5, 3, 6}}
	return pacificAtlantic(heights)
}
