package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
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

func RunCourseSchedule() bool {
	numCourses := 3
	prequisites := [][]int{{0, 1}, {2, 1}}
	return canFinish(numCourses, prequisites)
}
