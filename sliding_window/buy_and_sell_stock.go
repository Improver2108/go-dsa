package slidingwindow

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0
	for r < len(prices) {
		if prices[r] >= prices[l] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r++
	}
	return maxP
}

func RunMaxProfit() int {
	prices := []int{10, 1, 7, 5, 2}
	return maxProfit(prices)
}
