package dp1d

// Given an integer array nums, find a subarray that has the largest product, and return the product.

// A subarray is a contiguous non-empty sequence of elements within an array.

// You can assume the output will fit into a 32-bit integer.

// Note that the product of an array with a single element is the value of that element.

// Example 1:

// Input: nums = [2,4,-3,5]

// Output: 8
// Explanation: [2,4] has the largest product 8.

// Example 2:

// Input: nums = [-3,0,-2]

// Output: 0
// Explanation: The result cannot be 6, because [-3,-2] is not a subarray.

// Constraints:

// 1 <= nums.length <= 1000
// -10 <= nums[i] <= 10
// The product of any subarray of nums is guaranteed to fit in a 32-bit integer.

func maxProduct(nums []int) int {
	res, minEnding, maxEnding := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		temp := minEnding
		minEnding = min(maxEnding*nums[i], nums[i], temp*nums[i])
		maxEnding = max(maxEnding*nums[i], nums[i], temp*nums[i])
		res = max(res, maxEnding)
	}
	return res
}

func RunMaxProduct() int {
	nums := []int{-3, 0, -2, -6, -12}
	return maxProduct(nums)
}

// // [-3,0,-2,-6,-12]
// var maxNum = 0
// var minNum = -3
// var res = -3
