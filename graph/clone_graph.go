package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
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
