package slidingwindow

func totalFruit(fruits []int) int {
	start := 0
	basket := make(map[int]int)
	totalMax := 1
	for end := range fruits {
		basket[fruits[end]]++
		for len(basket) > 2 {
			basket[fruits[start]]--
			if basket[fruits[start]] == 0 {
				delete(basket, fruits[start])
			}
			start++
		}
		totalMax = max(totalMax, end-start+1)
	}
	return totalMax
}

func RunTotalFruit() int {
	fruits := []int{1, 0, 1, 4, 1, 4, 1, 2, 3}
	return totalFruit(fruits)
}
