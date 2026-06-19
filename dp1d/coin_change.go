package dp1d

import (
	"fmt"
	"math"
)

// You are given an integer array coins representing coins of different denominations (e.g. 1 dollar, 5 dollars, etc) and an integer amount representing a target amount of money.

// Return the fewest number of coins that you need to make up the exact target amount. If it is impossible to make up the amount, return -1.

// You may assume that you have an unlimited number of each coin.

// Example 1:

// Input: coins = [1,5,10], amount = 12

// Output: 3
// Explanation: 12 = 10 + 1 + 1. Note that we do not have to use every kind coin available.

func coinChangeMemoization(coins []int, amount int) int {
	var dfs func(amount int) int
	memo := make(map[int]int)
	memo[0] = 0
	dfs = func(amount int) int {
		if val, exists := memo[amount]; exists {
			return val
		}
		res := math.MaxInt32
		for _, amt := range coins {
			if amount-amt >= 0 {
				res = min(dfs(amount-amt)+1, res)
			}
		}
		memo[amount] = res
		return res
	}
	if dfs(amount) < math.MaxInt32 {
		return dfs(amount)
	}
	return -1
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for amt := 1; amt < amount+1; amt++ {
		for _, coin := range coins {
			if amt-coin >= 0 {
				dp[amt] = min(dp[amt-coin]+1, dp[amt])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func RunCoinChange() int {
	coins := []int{1, 2, 5}
	amount := 100
	fmt.Println(coins)
	return coinChange(coins, amount)
}
