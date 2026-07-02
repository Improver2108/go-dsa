package dp1d

// You are given an array of positive integers nums.

// Return true if you can partition the array into two subsets, subset1 and subset2 where sum(subset1) == sum(subset2). Otherwise, return false.

// Example 1:

// Input: nums = [1,2,3,4]

// Output: true
// Explanation: The array can be partitioned as [1, 4] and [2, 3].

// Example 2:

// Input: nums = [1,2,3,4,5]

// Output: false
// Constraints:

// 1 <= nums.length <= 100
// 1 <= nums[i] <= 50

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	memo := make(map[[2]int]bool)
	var isPossible func(i, n int) bool
	isPossible = func(i, n int) bool {
		if val, ok := memo[[2]int{i, n}]; ok {
			return val
		}
		if n == 0 {
			return true
		}
		for j := i; j < len(nums)-1; j++ {
			if n-nums[j] >= 0 {
				if isPossible(j+1, n-nums[j]) {
					memo[[2]int{i, n}] = true
					return true
				}
			}
		}
		memo[[2]int{i, n}] = false
		return false

	}
	return isPossible(0, sum/2)
}

func RunCanPartition() bool {
	nums := []int{1, 2, 3}
	return canPartition(nums)
}
