package dp2d

func changeMemo(amount int, coins []int) int {
	var dfs func(amt, i int) int
	memo := make([][]int, amount+1)
	for i := range memo {
		memo[i] = make([]int, len(coins)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	dfs = func(amt, i int) int {
		if amt == 0 {
			return 1
		}
		if i >= len(coins) || amt < 0 {
			return 0
		}
		if memo[amt][i] != -1 {
			return memo[amt][i]
		}
		res := dfs(amt-coins[i], i) + dfs(amt, i+1)
		memo[amt][i] = res
		return res

	}
	return dfs(amount, 0)
}

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, amount+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := range n + 1 {
		dp[0][i] = 1
	}

	for amt := range amount + 1 {
		for i := n - 1; i >= 0; i-- {
			dp[amt][i] = dp[amt][i+1]
			if amt-coins[i] >= 0 {
				dp[amt][i] += dp[amt-coins[i]][i]
			}
		}
	}
	return dp[amount][0]
}

func RunCoinChange() int {
	amount := 4
	coins := []int{1, 2, 3}
	return change(amount, coins)
}
