package dp2d

// You are given an array of integers nums and an integer target.

// For each number in the array, you can choose to either add or subtract it to a total sum.

// For example, if nums = [1, 2], one possible sum would be "+1-2=-1".
// If nums=[1,1], there are two different ways to sum the input numbers to get a sum of 0: "+1-1" and "-1+1".

// Return the number of different ways that you can build the expression such that the total sum equals target.

// Example 1:

// Input: nums = [2,2,2], target = 2

// Output: 3
// Explanation: There are 3 different ways to sum the input numbers to get a sum of 2.
// +2 +2 -2 = 2
// +2 -2 +2 = 2
// -2 +2 +2 = 2

// Constraints:

// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// -1000 <= target <= 1000

func findTargetSumWaysMemo(nums []int, target int) int {
	n := len(nums)
	var dfs func(total, i int) int
	memo := make(map[[2]int]int)
	dfs = func(total, i int) int {
		if total == target && i == n {
			return 1
		}
		if i >= len(nums) {
			return 0
		}
		if val, ok := memo[[2]int{total, i}]; ok {
			return val
		}
		memo[[2]int{total, i}] = dfs(total-nums[i], i+1) + dfs(total+nums[i], i+1)
		return memo[[2]int{total, i}]
	}
	return dfs(0, 0)
}

func findTargetSumWays(nums []int, target int) int {
	n, totalSum := len(nums), 0
	for _, num := range nums {
		totalSum += num
	}
	if target > totalSum || target < -totalSum {
		return 0
	}

	offset := totalSum

	dp := make([][]int, 2*totalSum+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[target+offset][n] = 1
	for i := n - 1; i >= 0; i-- {
		for total := -totalSum; total <= totalSum; total++ {
			ways := 0
			next := total - nums[i] + offset
			if next >= 0 && next <= 2*totalSum {
				ways += dp[next][i+1]
			}
			next = total + nums[i] + offset
			if next >= 0 && next <= 2*totalSum {
				ways += dp[next][i+1]
			}
			dp[total+offset][i] = ways
		}
	}
	return dp[offset][0]
}

func RunFindTargetSumWays() int {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	return findTargetSumWays(nums, target)
}
