package solution

// We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1.
// Return the same tree where every subtree (of the given tree) not containing a 1 has been removed.
// Recall that the subtree of a node X is X, plus every node that is a descendant of X.
//
// Examples:
//
// Input: [1,null,0,0,1]
// Output: [1,null,0,null,1]
//
// Input: [1,0,1,0,0,0,1]
// Output: [1,null,1,null,1]
//
// Input: [1,1,0,1,1,0,1,0]
// Output: [1,1,0,1,1,null,1]

func pruneTree(root *TreeNode) *TreeNode {
	if pruneTreeLoop(root) {
		return root
	}
	return nil
}

func pruneTreeLoop(node *TreeNode) bool {
	if node == nil {
		return false
	}
	l := pruneTreeLoop(node.Left)
	r := pruneTreeLoop(node.Right)
	if !l {
		node.Left = nil
	}
	if !r {
		node.Right = nil
	}
	return node.Val == 1 || l || r
}
