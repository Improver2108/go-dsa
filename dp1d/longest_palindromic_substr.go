package dp1d

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
