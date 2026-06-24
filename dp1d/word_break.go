package dp1d

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of dictionary words.

// You are allowed to reuse words in the dictionary an unlimited number of times. You may assume all dictionary words are unique.

// Example 1:

// Input: s = "neetcode", wordDict = ["neet","code"]

// Output: true
// Explanation: Return true because "neetcode" can be split into "neet" and "code".

// Example 2:

// Input: s = "applepenapple", wordDict = ["apple","pen","ape"]

// Output: true
// Explanation: Return true because "applepenapple" can be split into "apple", "pen" and "apple". Notice that we can reuse words and also not use all the words.

// Example 3:

// Input: s = "catsincars", wordDict = ["cats","cat","sin","in","car"]

// Output: false
// Constraints:

// 1 <= s.length <= 200
// 1 <= wordDict.length <= 100
// 1 <= wordDict[i].length <= 20
// s and wordDict[i] consist of only lowercase English letters.

func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]struct{})
	memo := make(map[int]bool)
	for _, word := range wordDict {
		set[word] = struct{}{}
	}
	var helper func(i int) bool
	memo[len(s)] = true
	helper = func(i int) bool {
		if val, ok := memo[i]; ok {
			return val
		}

		for j := i; j < len(s); j++ {
			if _, ok := set[s[i:j+1]]; ok {
				if helper(j + 1) {
					memo[i] = true
					return true
				}
			}
		}
		memo[i] = false
		return false
	}
	return helper(0)
}

func RunWordBreak() bool {
	s := "catsincars"
	wordDict := []string{"cats", "cat", "sin", "in", "car"}
	return wordBreak(s, wordDict)
}
