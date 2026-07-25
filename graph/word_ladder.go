package graph

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
	wordList := []string{"bat", "bag", "sag", "dag", "dot",}
	return ladderLength(beginWord, endWord, wordList)
}
