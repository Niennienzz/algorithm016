package solution

// Given the root of a binary tree, return the inorder traversal of its nodes' values.
// Recursive solution is trivial, could you do it iteratively?

// Time complexity: O(N), because the recursive function is T(N) = 2⋅T(N/2)+1.
// Space complexity: The worst case space required is O(N),
// and in the average case it's O(logN) where N is number of nodes.
func inorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0)
	inorderTraversalLoop(root, &ans)
	return ans
}

func inorderTraversalLoop(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	inorderTraversalLoop(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorderTraversalLoop(root.Right, ans)
}

// Time complexity: O(N).
// Space complexity: O(N).
func inorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		ans   = make([]int, 0)
		curr  = root
		stack = newTreeNodeStack()
	)

	for curr != nil || !stack.isEmpty() {
		for curr != nil {
			stack.push(curr)
			curr = curr.Left
		}
		curr = stack.pop()
		ans = append(ans, curr.Val)
		curr = curr.Right
	}

	return ans
}
