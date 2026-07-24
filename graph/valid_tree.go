package graph

func validTree(n int, edges [][]int) bool {
	if n == 1 && len(edges) == 0 {
		return true
	}
	finished := 0
	visited := make(map[int]bool)
	adjancencyList := make(map[int][]int)
	for i := range n {
		adjancencyList[i] = []int{}
	}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjancencyList[u] = append(adjancencyList[u], v)
		adjancencyList[v] = append(adjancencyList[v], u)
	}
	var isCycle func(currNode, prevNode int) bool
	isCycle = func(currNode, prevNode int) bool {
		if visited[currNode] {
			return true
		}
		if len(adjancencyList[currNode]) == 0 {
			return false
		}
		finished++
		visited[currNode] = true
		for _, node := range adjancencyList[currNode] {
			if node != prevNode && isCycle(node, currNode) {
				return true
			}
		}
		visited[currNode] = false
		adjancencyList[currNode] = []int{}
		return false
	}
	if isCycle(0, -1) {
		return false
	}
	return finished == n
}

func RunValidTree() bool {
	n := 1
	// edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	// edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}}
	// edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	edges := [][]int{}
	return validTree(n, edges)
}
