package graph

import "github.com/emirpasic/gods/v2/queues/arrayqueue"

func canFinishDetectCycle(numCourses int, prerequisites [][]int) bool {
	adjancencyList := make(map[int][]int)
	for i := range numCourses {
		adjancencyList[i] = []int{}
	}
	for _, prereq := range prerequisites {
		crs, pre := prereq[0], prereq[1]
		adjancencyList[crs] = append(adjancencyList[crs], pre)
	}
	visiting := make(map[int]bool)
	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if visiting[crs] {
			return false
		}
		if len(adjancencyList[crs]) == 0 {
			return true
		}
		visiting[crs] = true
		for _, ad := range adjancencyList[crs] {
			if !dfs(ad) {
				return false
			}
		}
		visiting[crs] = false
		adjancencyList[crs] = []int{}
		return true
	}
	for crs := range numCourses {
		if !dfs(crs) {
			return false
		}
	}
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjancencyList := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, preq := range prerequisites {
		src, dest := preq[0], preq[1]
		if _, ok := adjancencyList[src]; !ok {
			adjancencyList[src] = []int{}
		}
		adjancencyList[src] = append(adjancencyList[src], dest)
		indegree[dest]++
	}
	queue := arrayqueue.New[int]()
	for i, ind := range indegree {
		if ind == 0 {
			queue.Enqueue(i)
		}
	}
	finish := 0
	for !queue.Empty() {
		node, _ := queue.Dequeue()
		finish++
		for _, dest := range adjancencyList[node] {
			indegree[dest]--
			if indegree[dest] == 0 {
				queue.Enqueue(dest)
			}
		}
	}
	return finish == numCourses
}

func RunCourseSchedule() bool {
	numCourses := 3
	prequisites := [][]int{{0, 1}, {2, 1}}
	return canFinish(numCourses, prequisites)
}
