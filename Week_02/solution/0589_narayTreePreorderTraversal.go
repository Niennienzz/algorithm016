package solution

// Given an n-ary tree, return the preorder traversal of its nodes' values.

func preorder(root *Node) []int {
	ans := make([]int, 0)
	preorderLoop(root, &ans)
	return ans
}

func preorderLoop(node *Node, ans *[]int) {
	if node == nil {
		return
	}
	*ans = append(*ans, node.Val)
	for _, v := range node.Children {
		preorderLoop(v, ans)
	}
}
