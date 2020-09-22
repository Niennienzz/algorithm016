package solution

import (
	"math"
)

// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
//
// Examples:
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// Input: [3,9,20,null,null,15,7]
// Output: minimum depth = 2

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := math.MaxInt32
	if root.Left != nil {
		min = minInt(minDepth(root.Left), min)
	}
	if root.Right != nil {
		min = minInt(minDepth(root.Right), min)
	}

	return min + 1
}
