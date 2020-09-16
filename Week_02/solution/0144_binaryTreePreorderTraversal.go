package solution

// Given the root of a binary tree, return the preorder traversal of its nodes' values.
// Recursive solution is trivial, could you do it iteratively?

func preorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0)
	preorderTraversalLoop(root, &ans)
	return ans
}

func preorderTraversalLoop(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	preorderTraversalLoop(root.Left, ans)
	preorderTraversalLoop(root.Right, ans)
}

// Time complexity: O(N).
// Space complexity: O(N).
func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		ans   = make([]int, 0)
		curr  = root
		stack = newTreeNodeStack()
	)

	stack.push(curr)

	for !stack.isEmpty() {
		curr = stack.pop()
		ans = append(ans, curr.Val)
		if curr.Right != nil {
			stack.push(curr.Right)
		}
		if curr.Left != nil {
			stack.push(curr.Left)
		}
	}

	return ans
}
