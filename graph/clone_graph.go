package graph

import "github.com/emirpasic/gods/v2/queues/arrayqueue"

// Given a node in a connected undirected graph, return a deep copy of the graph.

// Each node in the graph contains an integer value and a list of its neighbors.

// class Node {
//     public int val;
//     public List<Node> neighbors;
// }
// The graph is shown in the test cases as an adjacency list. An adjacency list is a mapping of nodes to lists, used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

// For simplicity, nodes values are numbered from 1 to n, where n is the total number of nodes in the graph. The index of each node within the adjacency list is the same as the node's value (1-indexed).

// The input node will always be the first node in the graph and have 1 as the value.

// Example 1:

// Input: adjList = [[2],[1,3],[2]]

// Output: [[2],[1,3],[2]]
// Explanation: There are 3 nodes in the graph.
// Node 1: val = 1 and neighbors = [2].
// Node 2: val = 2 and neighbors = [1, 3].
// Node 3: val = 3 and neighbors = [2].

// Example 2:

// Input: adjList = [[]]

// Output: [[]]
// Explanation: The graph has one node with no neighbors.

// Example 3:

// Input: adjList = []

// Output: []
// Explanation: The graph is empty.

// Constraints:

// 0 <= The number of nodes in the graph <= 100.
// 1 <= Node.val <= 100
// There are no duplicate edges and no self-loops in the graph.

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
