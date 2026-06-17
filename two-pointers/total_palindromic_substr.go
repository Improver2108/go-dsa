package twopointers

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
	count := 0
	n := len(s)

	inner := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			count++
			l--
			r++
		}
	}

	for i := range n {
		inner(i, i)
		inner(i, i+1)
	}
	return count
}

func RunCountSubstrings() int {
	s := "aaabaaa"
	return countSubstrings(s)
}
