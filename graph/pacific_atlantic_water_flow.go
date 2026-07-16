package graph

import (
	"maps"
	"slices"
)

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
