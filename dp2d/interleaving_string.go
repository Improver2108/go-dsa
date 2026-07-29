package dp2d

// You are given three strings s1, s2, and s3. Return true if s3 is formed by interleaving s1 and s2 together or false otherwise.

// Interleaving two strings s and t is done by dividing s and t into n and m substrings respectively, where the following conditions are met

// |n - m| <= 1, i.e. the difference between the number of substrings of s and t is at most 1.
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// Interleaving s and t is s1 + t1 + s2 + t2 + ... or t1 + s1 + t2 + s2 + ...
// You may assume that s1, s2 and s3 consist of lowercase English letters.

// Input: s1 = "aaaa", s2 = "bbbb", s3 = "aabbbbaa"

// Output: true
// Explanation: We can split s1 into ["aa", "aa"], s2 can remain as "bbbb" and s3 is formed by interleaving ["aa", "aa"] and "bbbb".

func isInterleave(s1 string, s2 string, s3 string) bool {
	var dfs func(i, j, k int) bool
	n1, n2, n3 := len(s1), len(s2), len(s3)
	memo := make(map[[2]int]bool)
	dfs = func(i, j, k int) bool {
		if k == n3 {
			return i == n1 && j == n2
		}
		if value, ok := memo[[2]int{i, j}]; ok {
			return value
		}
		if i < n1 && s1[i] == s3[k] {
			if dfs(i+1, j, k+1) {
				memo[[2]int{i, j}] = true
				return true
			}
		}
		if j < n2 && s2[j] == s3[k] {
			if dfs(i, j+1, k+1) {
				memo[[2]int{i, j}] = true
				return true
			}
		}
		memo[[2]int{i, j}] = false
		return false
	}
	return dfs(0, 0, 0)
}

func RunIsInterleave() bool {
	s1 := "aaaa"
	s2 := "bbbb"
	s3 := "aabbbbaa"
	return isInterleave(s1, s2, s3)
}
