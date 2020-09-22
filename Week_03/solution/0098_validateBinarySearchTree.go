package solution

// Given a binary tree, determine if it is a valid binary search tree (BST).
// Assume a BST is defined as follows:
//  - The left subtree of a node contains only nodes with keys less than the node's key.
//  - The right subtree of a node contains only nodes with keys greater than the node's key.
//  - Both the left and right subtrees must also be binary search trees.
//
// Examples:
//
//    2
//   / \
//  1   3
// Input: [2,1,3]
// Output: true
//
//    5
//   / \
//  1   4
//     / \
//    3   6
//
// Input: [5,1,4,null,null,3,6]
// Output: false
//
//    10
//   /  \
//  5    15
//      / \
//     6   20
//
// Input: [10,5,15,null,null,6,20]
// Output: false

func isValidBST(root *TreeNode) bool {
	return isValidBSTLoop(root, nil, nil)
}

func isValidBSTLoop(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}

	val := node.Val
	if lower != nil && val <= *lower {
		return false
	}
	if upper != nil && val >= *upper {
		return false
	}

	if !isValidBSTLoop(node.Right, &val, upper) {
		return false
	}
	if !isValidBSTLoop(node.Left, lower, &val) {
		return false
	}

	return true
}
