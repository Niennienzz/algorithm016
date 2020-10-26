package solution

// Given the root of a binary tree, return the tilt of the whole tree.
// The tilt of a tree node is the absolute difference between the sum
// of all left subtree node values and all right subtree node values.
// Null node has tilt 0.
// The tilt of the whole tree is the sum of all nodes' tilt.
//
// Examples:
//
// Input: root = [1,2,3]
// Output: 1
// Explanation:
//  - Tilt of node 2 : 0
//  - Tilt of node 3 : 0
//  - Tilt of node 1 : |2-3| = 1
//  - Tilt of binary tree : 0 + 0 + 1 = 1
//
// Input: root = [4,2,9,3,5,null,7]
// Output: 15
// Explanation:
//  - Tilt of node 3 : 0
//  - Tilt of node 5 : 0
//  - Tilt of node 7 : 0
//  - Tilt of node 2 : |3-5| = 2
//  - Tilt of node 9 : |0-7| = 7
//  - Tilt of node 4 : |(3+5+2)-(9+7)| = 6
//  - Tilt of binary tree : 0 + 0 + 0 + 2 + 7 + 6 = 15

func findTilt(root *TreeNode) int {
	var ans int
	findTiltLoop(root, &ans)
	return ans
}

func findTiltLoop(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	l := findTiltLoop(node.Left, ans)
	r := findTiltLoop(node.Right, ans)
	*ans += absInt(l - r)
	return l + r + node.Val
}
