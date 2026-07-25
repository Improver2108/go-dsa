package graph

import (
	"slices"
)

func findRedundantConnection1(edges [][]int) []int {
	adj := make(map[int][]int)
	visited := make(map[int]bool)
	var dfs func(curr, prev int) bool
	dfs = func(curr, prev int) bool {
		if visited[curr] {
			return true
		}
		visited[curr] = true
		for _, nei := range adj[curr] {
			if nei == prev {
				continue
			}
			if dfs(nei, curr) {
				return true
			}
		}
		return false
	}
	for _, pair := range edges {
		u, v := pair[0], pair[1]
		adj[v] = append(adj[v], u)
		adj[u] = append(adj[u], v)
		for i := range len(edges) + 1 {
			visited[i] = false
		}
		if dfs(u, -1) {
			return []int{u, v}
		}
	}

	return []int{}
}

func findRedundantConnection(edges [][]int) []int {
	adj := make(map[int][]int)
	visit := make(map[int]bool)
	cycle := make(map[int]bool)
	cycleStart := -1
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	var dfs func(curr, prev int) bool
	dfs = func(curr, prev int) bool {
		if visit[curr] {
			cycleStart = curr
			return true
		}
		visit[curr] = true
		for _, neig := range adj[curr] {
			if neig == prev {
				continue
			}
			if dfs(neig, curr) {
				if cycleStart != -1 {
					cycle[curr] = true
				}
				if cycleStart == curr {
					cycleStart = -1
				}
				return true
			}
		}
		return false
	}
	dfs(1, -1)
	for _, edge := range slices.Backward(edges) {
		u, v := edge[0], edge[1]
		if cycle[u] && cycle[v] {
			return []int{u, v}
		}
	}
	return []int{}
}

func RunFindRedundantConnection() []int {
	edges := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
	return findRedundantConnection(edges)
}
