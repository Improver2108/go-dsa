package graph

import "github.com/emirpasic/gods/v2/queues/arrayqueue"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraphDfs(node *Node) *Node {
	mapOldToNew := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if val, ok := mapOldToNew[node]; ok {
			return val
		}
		copy := &Node{Val: node.Val}
		mapOldToNew[node] = copy
		for _, nei := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors, dfs(nei))
		}
		return copy
	}
	return dfs(node)
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	oldToNew := make(map[*Node]*Node)
	oldToNew[node] = &Node{Val: node.Val}
	queue := arrayqueue.New[*Node]()
	queue.Enqueue(node)
	for !queue.Empty() {
		curr, _ := queue.Dequeue()
		for _, nei := range curr.Neighbors {
			if _, ok := oldToNew[nei]; !ok {
				oldToNew[nei] = &Node{Val: nei.Val}
				queue.Enqueue(nei)
			}
			oldToNew[curr].Neighbors = append(oldToNew[curr].Neighbors, oldToNew[nei])
		}
	}
	return oldToNew[node]
}
