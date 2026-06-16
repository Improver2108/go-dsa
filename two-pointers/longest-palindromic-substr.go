package twopointers

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
	resIndex, resLen := 0, 0
	n := len(s)

	inner := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > resLen {
				resIndex = l
				resLen = r - l + 1
			}
			l--
			r++
		}
	}

	for i := range n {
		inner(i, i)
		inner(i, i+1)
	}
	return s[resIndex : resIndex+resLen]
}

func LongestPalindrome() string {
	s := "ababd"
	return longestPalindrome(s)
}
