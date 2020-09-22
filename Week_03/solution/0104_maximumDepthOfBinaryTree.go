package solution

// Given a binary tree, find its maximum depth.
// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
// Examples:
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// Input: [3,9,20,null,null,15,7]
// Output: maximum depth = 3

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	a := maxDepth(root.Left) + 1
	b := maxDepth(root.Right) + 1

	return maxInt(a, b)
}
