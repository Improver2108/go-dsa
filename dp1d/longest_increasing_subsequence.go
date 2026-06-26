package dp1d

// Given an integer array nums, return the length of the longest strictly increasing subsequence.

// A subsequence is a sequence that can be derived from the given sequence by deleting some or no elements without changing the relative order of the remaining characters.

// For example, "cat" is a subsequence of "crabt".
// Example 1:

// Input: nums = [9,1,4,2,3,3,7]

// Output: 4
// Explanation: The longest increasing subsequence is [1,2,3,7], which has a length of 4.

// Example 2:

// Input: nums = [0,3,1,3,2,3]

// Output: 4
// Constraints:

// 1 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000

func lengthOfLIS(nums []int) int {
	stack := []int{nums[0]}
	res := len(stack)
	for i := 1; i < len(nums); i++ {
		curr := len(stack) - 1
		for curr >= 0 && nums[i] <= stack[curr] {
			curr--
		}
		if curr < len(stack)-1 {
			stack[curr+1] = nums[i]
		} else {
			stack = append(stack, nums[i])
			res = max(res, len(stack))
		}
	}
	return res
}

func RunLengthOfLIS() int {
	nums := []int{9, 1, 4, 2, 3, 3, 7}
	return lengthOfLIS(nums)
}
