package graph

import "github.com/emirpasic/gods/v2/queues/arrayqueue"

// You are given a
// m
// ×
// n
// m×n 2D grid initialized with these three possible values:

// -1 - A water cell that can not be traversed.
// 0 - A treasure chest.
// INF - A land cell that can be traversed. We use the integer 2^31 - 1 = 2147483647 to represent INF.
// Fill each land cell with the distance to its nearest treasure chest. If a land cell cannot reach a treasure chest then the value should remain INF.

// Assume the grid can only be traversed up, down, left, or right.

// Modify the grid in-place.

// Example 1:

// Input: [
//   [2147483647,-1,0,2147483647],
//   [2147483647,2147483647,2147483647,-1],
//   [2147483647,-1,2147483647,-1],
//   [0,-1,2147483647,2147483647]
// ]

// Output: [
//   [3,-1,0,1],
//   [2,2,1,-1],
//   [1,-1,2,-1],
//   [0,-1,3,4]
// ]
// Example 2:

// Input: [
//   [0,-1],
//   [2147483647,2147483647]
// ]

// Output: [
//   [0,-1],
//   [1,2]
// ]
// Constraints:

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 100
// grid[i][j] is one of {-1, 0, 2147483647}

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
