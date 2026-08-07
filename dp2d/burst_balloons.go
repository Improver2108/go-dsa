package dp2d

func maxCoinsMemo(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}
		if dp[l][r] > 0 {
			return dp[l][r]
		}
		dp[l][r] = 0
		for i := l; i <= r; i++ {
			coin := nums[l-1] * nums[i] * nums[r+1]
			coin += dfs(l, i-1) + dfs(i+1, r)
			dp[l][r] = max(dp[l][r], coin)
		}
		return dp[l][r]
	}
	return dfs(1, len(nums)-2)
}

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for l := n; l >= 1; l-- {
		for r := l; r <= n; r++ {
			for i := l; i <= r; i++ {
				coins := nums[l-1] * nums[i] * nums[r+1]
				coins += dp[l][i-1] + dp[i+1][r]
				dp[l][r] = max(dp[l][r], coins)
			}
		}
	}
	return dp[1][n]
}

func RunMaxCoins() int {
	nums := []int{4, 2, 3, 7}
	return maxCoins(nums)
}
