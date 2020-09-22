package solution

// Invert a binary tree.
//
// Examples:
//
// Input:
//     4
//   /   \
//  2     7
// / \   / \
// 1   3 6   9
//
// Output: [[""]]
//     4
//   /   \
//  7     2
// / \   / \
// 9   6 3   1

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
