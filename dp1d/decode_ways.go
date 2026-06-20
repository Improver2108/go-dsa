package dp1d

// A string consisting of uppercase english characters can be encoded to a number using the following mapping:

// 'A' -> "1"
// 'B' -> "2"
// ...
// 'Z' -> "26"
// To decode a message, digits must be grouped and then mapped back into letters using the reverse of the mapping above. There may be multiple ways to decode a message. For example, "1012" can be mapped into:

// "JAB" with the grouping (10 1 2)
// "JL" with the grouping (10 12)
// The grouping (1 01 2) is invalid because 01 cannot be mapped into a letter since it contains a leading zero.

// Given a string s containing only digits, return the number of ways to decode it. You can assume that the answer fits in a 32-bit integer.

func numDecodingsMemo(s string) int {
	var dfs func(int) int
	memo := make(map[int]int)
	dfs = func(i int) int {
		if i == len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}

		if val, ok := memo[i]; ok {
			return val
		}
		res := dfs(i + 1)
		if i < len(s)-1 {
			if (s[i] == '1') || (s[i] == '2' && s[i+1] <= '6') {
				res += dfs(i + 2)
			}
		}
		memo[i] = res
		return res
	}
	return dfs(0)
}

func numDecodings(s string) int {
	dp := make(map[int]int)
	dp[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
			if i < len(s)-1 {
				if (s[i] == '1') || (s[i] == '2' && s[i+1] <= '6') {
					dp[i] += dp[i+2]
				}
			}
		}
	}
	return dp[0]
}

//"2632"

func RunNumDecodings() int {
	s := "1012"
	return numDecodingsMemo(s)
}
