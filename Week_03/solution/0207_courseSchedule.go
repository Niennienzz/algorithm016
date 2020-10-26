package solution

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
// Some courses may have prerequisites, for example:
// To take course 0 you have to first take course 1, which is expressed as a pair: [0,1]
// Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?
//
// Examples:
//
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation:
//  - There are a total of 2 courses to take.
//  - To take course 1 you should have finished course 0. So it is possible.
//
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation:
//  - There are a total of 2 courses to take.
//  - To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1.
//  - So it is impossible.

// 拓扑排序.
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	// Build the graph.
	graph := make(map[int]*graphNode)
	for _, relation := range prerequisites {
		// [0,1]: To take course 0 you have to first take course 1.
		prevCourse := getOrCreateGraphNode(graph, relation[1])
		nextCourse := getOrCreateGraphNode(graph, relation[0])
		prevCourse.addOutNode(relation[0])
		nextCourse.inDegrees += 1
	}

	var (
		totalDeps    = len(prerequisites)
		removedEdges = 0
		noDepCourses = newIntQueue()
	)

	// Start from courses without prerequisites.
	for course := 0; course < numCourses; course++ {
		if courseNode, ok := graph[course]; ok && courseNode.inDegrees == 0 {
			noDepCourses.push(course)
		}
	}

	for !noDepCourses.isEmpty() {
		course := noDepCourses.pop()
		for _, nextCourse := range graph[course].outNodes {
			nextCourseNode := graph[nextCourse]
			nextCourseNode.inDegrees -= 1
			removedEdges += 1
			if nextCourseNode.inDegrees == 0 {
				noDepCourses.push(nextCourse)
			}
		}
	}

	// If there are still some edges left, then there exist some cycles.
	// Due to the dead-lock (dependencies), we cannot remove the cyclic edges.
	return removedEdges == totalDeps
}

func getOrCreateGraphNode(graph map[int]*graphNode, course int) *graphNode {
	if node, ok := graph[course]; ok {
		return node
	}
	node := &graphNode{}
	graph[course] = node
	return node
}

type graphNode struct {
	inDegrees int
	outNodes  []int
}

func (node *graphNode) addOutNode(v int) {
	node.outNodes = append(node.outNodes, v)
}
