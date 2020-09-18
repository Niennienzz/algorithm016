package solution

// Given an n-ary tree, return the level order traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal,
// each group of children is separated by the null value (See examples).
//
// Examples:
//
// Input: root = [1,null,3,2,4,null,5,6]
//      1
//     /|\
//    3 2 4
//   / \
//  5   6
// Output: [[1],[3,2,4],[5,6]]

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var (
		ans   = make([][]int, 0)
		queue = newNodeQueue()
	)

	queue.push(root)

	for !queue.isEmpty() {
		// Get current level values, and append all children
		// of all nodes into the queue as they are the next level.
		var (
			size = queue.size()
			curr = make([]int, size)
		)

		for i := 0; i < size; i++ {
			node := queue.pop()
			curr[i] = node.Val
			queue.pushAll(node.Children)
		}

		ans = append(ans, curr)
	}

	return ans
}
