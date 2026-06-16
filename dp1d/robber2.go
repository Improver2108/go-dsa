package dp1d

// You are given an integer array nums where nums[i] represents the amount of money the ith house has. The houses are arranged in a circle, i.e. the first house and the last house are neighbors.

// You are planning to rob money from the houses, but you cannot rob two adjacent houses because the security system will automatically alert the police if two adjacent houses were both broken into.

// Return the maximum amount of money you can rob without alerting the police.

// Example 1:

// Input: nums = [3,4,3]

// Output: 4
// Explanation: You cannot rob nums[0] + nums[2] = 6 because nums[0] and nums[2] are adjacent houses. The maximum you can rob is nums[1] = 4.

// Example 2:

// Input: nums = [2,9,8,3,6]

// Output: 15
// Explanation: You cannot rob nums[0] + nums[2] + nums[4] = 16 because nums[0] and nums[4] are adjacent houses. The maximum you can rob is nums[1] + nums[4] = 15.

// Constraints:

// 1 <= nums.length <= 100
// 0 <= nums[i] <= 200

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
