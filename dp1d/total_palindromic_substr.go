package dp1d

// Given a string s, return the number of substrings within s that are palindromes.

// A palindrome is a string that reads the same forward and backward.

// Example 1:

// Input: s = "abc"

// Output: 3
// Explanation: "a", "b", "c".

// Example 2:

// Input: s = "aaa"

// Output: 6
// Explanation: "a", "a", "a", "aa", "aa", "aaa". Note that different substrings are counted as different palindromes even if the string contents are the same.

// Constraints:

// 1 <= s.length <= 1000
// s consists of lowercase English letters.

func countSubstrings(s string) int {
	n, count := len(s), 0
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, n)
	}
	for i := range n {
		for j := range i + 1 {
			if s[i] == s[j] && (i-j <= 2 || dp[i-1][j+1]) {
				dp[i][j] = true
				count++
			}
		}
	}
	return count
}

func RunCountSubstrings() int {
	s := "aaabaaa"
	return countSubstrings(s)
}
