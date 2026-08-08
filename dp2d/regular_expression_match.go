package dp2d

// You are given an input string s consisting of lowercase english letters, and a pattern p consisting of lowercase english letters, as well as '.', and '*' characters.

// Return true if the pattern matches the entire input string, otherwise return false.

// '.' Matches any single character
// '*' Matches zero or more of the preceding element.

// Example 1:

// Input: s = "aa", p = ".b"

// Output: false
// Explanation: Regardless of which character we choose for the '.' in the pattern, we cannot match the second character in the input string.

// Example 2:

// Input: s = "nnn", p = "n*"

// Output: true
// Explanation: '*' means zero or more of the preceding element, 'n'. We choose 'n' to repeat three times.

// Example 3:

// Input: s = "xyz", p = ".*z"

// Output: true
// Explanation: The pattern ".*" means zero or more of any character, so we choose ".." to match "xy" and "z" to match "z".

// Constraints:

// 1 <= s.length <= 20
// 1 <= p.length <= 20
// Each appearance of '*', will be preceded by a valid character or '.'

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if j >= n {
			return i >= m
		}
		if val, ok := dp[[2]int{i, j}]; ok {
			return val
		}
		match := i < m && (p[j] == '.' || s[i] == p[j])
		if j+1 < n && p[j+1] == '*' {
			dp[[2]int{i, j}] = (match && dfs(i+1, j)) || dfs(i, j+2)
			return dp[[2]int{i, j}]
		}
		if match {
			dp[[2]int{i, j}] = dfs(i+1, j+1)
			return dp[[2]int{i, j}]
		}
		dp[[2]int{i, j}] = false
		return dp[[2]int{i, j}]
	}
	return dfs(0, 0)
}

func RunRegularExpression() bool {
	s := "nnn"
	p := "n*"
	return isMatch(s, p)
}
