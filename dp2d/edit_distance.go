package dp2d

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m {
			return n - j
		}
		if j >= n {
			return m - i
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		res := 0
		if word1[i] == word2[j] {
			res = dfs(i+1, j+1)
			memo[i][j] = res
			return res
		}
		res = 1 + min(dfs(i+1, j), dfs(i+1, j+1), dfs(i, j+1))
		memo[i][j] = res
		return res
	}
	return dfs(0, 0)
}

func RunMinDistance() int {
	word1 := "neatcdee"
	word2 := "neetcode"
	return minDistance(word1, word2)
}
