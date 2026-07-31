package dp2d

// Longest Increasing Path in Matrix
// Hard
// Topics
// Company Tags
// Hints
// You are given a 2-D grid of integers matrix, where each integer is greater than or equal to 0.

// Return the length of the longest strictly increasing path within matrix.

// From each cell within the path, you can move either horizontally or vertically. You may not move diagonally.

// Example 1:
// Input: matrix = [[5,5,3],[2,3,6],[1,1,1]]

// Output: 4
//
// Explanation: The longest increasing path is [1, 2, 3, 6] or [1, 2, 3, 5].
// Input: matrix = [[1,2,3],[2,1,4],[7,6,5]]

// Output: 7
// Explanation: The longest increasing path is [1, 2, 3, 4, 5, 6, 7].

// Constraints:

// 1 <= matrix.length, matrix[i].length <= 100

func longestIncreasingPathDFS(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var dfs func(r, c, prev int) int
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dfs = func(r, c, prev int) int {
		if r < 0 || r >= m || c < 0 || c >= n || matrix[r][c] <= prev {
			return 0
		}
		if dp[r][c] != 0 {
			return dp[r][c]
		}
		res := 1
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			res = max(dfs(nr, nc, matrix[r][c])+1, res)
		}
		dp[r][c] = res
		return res
	}
	LIP := 1
	for i := range m {
		for j := range n {
			LIP = max(LIP, dfs(i, j, -1))
		}
	}
	return LIP
}

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	indegree := make([][]int, m)
	for i := range indegree {
		indegree[i] = make([]int, n)
	}
	for r := range m {
		for c := range n {
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n || matrix[nr][nc] >= matrix[r][c] {
					continue
				}
				indegree[r][c]++
			}
		}
	}
	q := [][]int{}
	for r := range m {
		for c := range n {
			if indegree[r][c] != 0 {
				continue
			}
			q = append(q, []int{r, c})
		}
	}
	lis := 0
	for len(q) > 0 {
		size := len(q)
		for range size {
			r, c := q[0][0], q[0][1]
			q = q[1:]
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n || matrix[nr][nc] <= matrix[r][c] {
					continue
				}
				indegree[nr][nc]--
				if indegree[nr][nc] == 0 {
					q = append(q, []int{nr, nc})
				}
			}
		}
		lis++
	}
	return lis
}

func RunLongestIncreasingPath() int {
	matrix := [][]int{{5, 5, 3}, {2, 3, 6}, {1, 1, 1}}
	return longestIncreasingPath(matrix)
}
