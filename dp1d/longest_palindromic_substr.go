package dp1d

// Given a string s, return the longest substring of s that is a palindrome.

// A palindrome is a string that reads the same forward and backward.

// If there are multiple palindromic substrings that have the same length, return any one of them.

// Example 1:

// Input: s = "ababd"

// Output: "bab"
// Explanation: Both "aba" and "bab" are valid answers.

// Example 2:

// Input: s = "abbc"

// Output: "bb"
// Constraints:

// 1 <= s.length <= 1000
// s contains only digits and English letters.

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	resIndex, resLen := 0, 0
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := range n {
		for j := range i + 1 {
			if s[i] == s[j] && (i-j <= 2 || dp[i-1][j+1]) {
				dp[i][j] = true
				if i-j+1 > resLen {
					resIndex = j
					resLen = i - j + 1
				}
			}
		}
	}
	return s[resIndex : resIndex+resLen]
}

// ababd

func LongestPalindrome() string {
	s := "ababd"
	return longestPalindrome(s)
}
