package graph

// You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:

// Connect: A cell is connected to adjacent cells horizontally or vertically.
// Region: To form a region connect every 'O' cell. Regions can have any shape; they do not need to be squares or rectangles.
// Surround: A region is surrounded if none of the 'O' cells in that region are on the edge of the board. Such regions are completely enclosed by 'X' cells.
// To capture a surrounded region, replace all 'O's with 'X's in-place within the original board. You do not need to return anything.

// Example 1:

// Input: board = [
//   ["X","X","X","X"],
//   ["X","O","O","X"],
//   ["X","X","O","X"],
//   ["X","O","X","X"]
// ]

// Output: [
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","X","X","X"],
//   ["X","O","X","X"]
// ]
// Explanation: The bottom 'O' region is not captured because it touches the edge of the board, so it cannot be surrounded.

// Example 2:

// Input: board = [["X"]]

// Output: [["X"]]
// Constraints:

// 1 <= board.length, board[i].length <= 200
// board[i][j] is 'X' or 'O'.

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	queue := [][]int{}
	for i := range m {
		for j := range n {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				queue = append(queue, []int{i, j})
			}
		}
	}

	directions := [...][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]
		board[row][col] = 'T'
		for _, dir := range directions {
			nr, nc := row+dir[0], col+dir[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || board[nr][nc] != 'O' {
				continue
			}
			board[nr][nc] = 'T'
			queue = append(queue, []int{nr, nc})
		}
	}
	for i := range m {
		for j := range n {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'T':
				board[i][j] = 'O'
			}
		}
	}
}

func RunSurroundedRegions() [][]byte {
	board := [][]byte{{'X', 'X', 'X'}, {'X', 'O', 'X'}, {'X', 'X', 'X'}}
	solve(board)
	return board
}
