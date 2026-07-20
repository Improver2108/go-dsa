package graph

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
