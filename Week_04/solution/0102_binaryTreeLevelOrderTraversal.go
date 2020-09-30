package solution

// Given a binary tree, return the level order traversal of its nodes' values.
//
// Examples:
//
// Input: [3,9,20,null,null,15,7]
//    3
//   / \
//  9  20
//    /  \
//   15   7
// Output: [[3],[9,20],[15,7]]

// Using BFS search.
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		ans   = make([][]int, 0)
		queue = newTreeNodeQueue()
	)
	queue.push(root)

	for !queue.isEmpty() {
		size := queue.len()
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.pop()
			level[i] = node.Val
			if node.Left != nil {
				queue.push(node.Left)
			}
			if node.Right != nil {
				queue.push(node.Right)
			}
		}
		ans = append(ans, level)
	}

	return ans
}
