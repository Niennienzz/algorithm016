package solution

import (
	"math"
)

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

// 递归写法: 按二叉搜索树的定义.
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

// 二叉搜索树的中序遍历为升序排列.
func isValidBSTInorder(root *TreeNode) bool {
	return newIsValidBSTInorderHelper().loop(root)
}

type isValidBSTInorderHelper struct {
	lastVal int
	flag    bool
}

func newIsValidBSTInorderHelper() *isValidBSTInorderHelper {
	return &isValidBSTInorderHelper{
		lastVal: math.MinInt64,
		flag:    true,
	}
}

func (x *isValidBSTInorderHelper) loop(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if x.flag && node.Left != nil {
		x.loop(node.Left)
	}

	if node.Val <= x.lastVal {
		x.flag = false
	}

	x.lastVal = node.Val

	if x.flag && node.Right != nil {
		x.loop(node.Right)
	}

	return x.flag
}
