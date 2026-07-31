package dp2d

type state struct {
	word string
	i    int
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	var dfs func(i, j int) int
	memo := make(map[[2]int]int)
	dfs = func(i, j int) int {
		if j == n {
			return 1
		}
		if i == m {
			return 0
		}
		if val, ok := memo[[2]int{i, j}]; ok {
			return val
		}
		res := dfs(i+1, j)
		if s[i] == t[j] {
			res += dfs(i+1, j+1)
		}
		memo[[2]int{i, j}] = res
		return res
	}
	return dfs(0, 0)
}

func RunNumDistinct() int {
	s := "caaat"
	t := "cat"
	return numDistinct(s, t)
}
