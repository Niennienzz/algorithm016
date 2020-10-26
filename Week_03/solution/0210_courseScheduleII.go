package solution

// 拓扑排序.
func findOrder(numCourses int, prerequisites [][]int) []int {
	// Build the graph.
	graph := make(map[int]*graphNode)
	for course := 0; course < numCourses; course++ {
		graph[course] = &graphNode{}
	}
	for _, relation := range prerequisites {
		// [0,1]: To take course 0 you have to first take course 1.
		prevCourse := graph[relation[1]]
		nextCourse := graph[relation[0]]
		prevCourse.addOutNode(relation[0])
		nextCourse.inDegrees += 1
	}

	var (
		totalDeps    = len(prerequisites)
		removedEdges = 0
		noDepCourses = newIntQueue()
		ans          []int
	)

	// Start from courses without prerequisites.
	for course := 0; course < numCourses; course++ {
		if courseNode, ok := graph[course]; ok && courseNode.inDegrees == 0 {
			noDepCourses.push(course)
		}
	}

	for !noDepCourses.isEmpty() {
		course := noDepCourses.pop()
		ans = append(ans, course)
		for _, nextCourse := range graph[course].outNodes {
			nextCourseNode := graph[nextCourse]
			nextCourseNode.inDegrees -= 1
			removedEdges += 1
			if nextCourseNode.inDegrees == 0 {
				noDepCourses.push(nextCourse)
			}
		}
	}

	if removedEdges == totalDeps {
		return ans
	}

	return nil
}
