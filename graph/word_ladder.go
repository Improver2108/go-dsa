package graph

// You are given two words, beginWord and endWord, and also a list of words wordList. All of the given words are of the same length, consisting of lowercase English letters, and are all distinct.

// Your goal is to transform beginWord into endWord by following the rules:

// You may transform beginWord to any word within wordList, provided that at exactly one position the words have a different character, and the rest of the positions have the same characters.
// You may repeat the previous step with the new word that you obtain, and you may do this as many times as needed.
// Return the minimum number of words within the transformation sequence needed to obtain the endWord, or 0 if no such sequence exists.

// Example 1:

// Input: beginWord = "cat", endWord = "sag", wordList = ["bat","bag","sag","dag","dot"]

// Output: 4
// Explanation: The transformation sequence is "cat" -> "bat" -> "bag" -> "sag".

// Example 2:

// Input: beginWord = "cat", endWord = "sag", wordList = ["bat","bag","sat","dag","dot"]

// Output: 0
// Explanation: There is no possible transformation sequence from "cat" to "sag" since the word "sag" is not in the wordList.

// Constraints:

// 1 <= beginWord.length <= 10
// 1 <= wordList.length <= 100

func ladderLength(beginWord, endWord string, wordList []string) int {
	q := []string{beginWord}
	visit := make(map[string]bool)
	visit[beginWord] = true
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	steps := 1
	for len(q) > 0 {
		size := len(q)
		for range size {
			word := q[0]
			q = q[1:]
			if word == endWord {
				return steps
			}
			chars := []byte(word)
			for i := range chars {
				original := chars[i]
				for ch := byte('a'); ch <= 'z'; ch++ {
					if original == ch {
						continue
					}
					chars[i] = ch
					next := string(chars)
					_, exists := wordSet[next]
					if !exists || visit[next] {
						continue
					}
					visit[next] = true
					q = append(q, next)
				}
				chars[i] = original
			}
		}
		steps++
	}
	return 0
}

func RunLadderLength() int {
	beginWord := "cat"
	endWord := "cat"
	wordList := []string{"bat", "bag", "sag", "dag", "dot"}
	return ladderLength(beginWord, endWord, wordList)
}
