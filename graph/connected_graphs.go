package graph

// Number of Connected Components in an Undirected Graph
// Medium
// Topics
// Company Tags
// Hints
// You have a graph of n nodes. You are given an integer n and an array edges where edges[i] = [aᵢ, bᵢ] indicates that there is an edge between aᵢ and bᵢ in the graph.

// Return the number of connected components in the graph.

// Example 1:

// Input:
// n = 5, edges = [[0,1],[1,2],[3,4]]

// Output: 2
// Example 2:

// Input:
// n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]

// Output: 1
// Constraints:

// 1 <= n <= 2000
// 1 <= edges.length <= 5000
// edges[i].length == 2
// 0 <= aᵢ <= bᵢ < n
// aᵢ != bᵢ
// There are no repeated edges.

func countComponents(n int, edges [][]int) int {
	visited := make(map[int]bool)
	adj := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	bfs := func(node int) {
		queue := []int{node}
		visited[node] = true
		for len(queue) > 0 {
			currNode := queue[0]
			queue = queue[1:]
			for _, neigh := range adj[currNode] {
				if visited[neigh] {
					continue
				}
				visited[neigh] = true
				queue = append(queue, neigh)
			}
		}
	}
	total := 0
	for i := range n {
		if !visited[i] {
			bfs(i)
			total++
		}
	}
	return total
}

func RunCountComponents() int {
	n := 5
	// edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	return countComponents(n, edges)
}
