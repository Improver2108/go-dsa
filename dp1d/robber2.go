package dp1d

func findMax(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	return max(findMax(nums[1:]), findMax(nums[:len(nums)-1]))
}

func Roberr2() int {
	nums := []int{2}
	return rob(nums)
}
