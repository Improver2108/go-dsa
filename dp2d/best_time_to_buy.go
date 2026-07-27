package dp2d

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 1; j >= 0; j-- {
			if j == 1 {
				dp[i][j] = max(dp[i+1][0]-prices[i], dp[i+1][1])
			} else {
				sell := prices[i]
				if i+2 < n {
					sell += dp[i+2][1]
				}
				dp[i][j] = max(sell, dp[i+1][0])
			}
		}
	}
	return dp[0][1]
}

func RunMaxProfit() int {
	prices := []int{1, 3, 4, 0, 4}
	return maxProfit(prices)
}
