package graph

import "github.com/emirpasic/gods/v2/queues/arrayqueue"

var INF = 2147483647

func islandsAndTreasure(grid [][]int) {
	m, n := len(grid), len(grid[0])
	q := arrayqueue.New[[2]int]()
	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				q.Enqueue([2]int{i, j})
			}
		}
	}
	if q.Empty() {
		return
	}
	directions := [...][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for !q.Empty() {
		node, _ := q.Dequeue()
		row, col := node[0], node[1]
		for _, dir := range directions {
			nr, nc := row+dir[0], col+dir[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] != INF {
				continue
			}
			grid[nr][nc] = grid[row][col] + 1
			q.Enqueue([2]int{nr, nc})
		}
	}

}

func RunIslandsAndTreasure() [][]int {
	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}
	islandsAndTreasure(rooms)
	return rooms
}

// Time taken: 1.483µs
