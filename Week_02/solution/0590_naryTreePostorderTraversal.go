package solution

// Given an n-ary tree, return the postorder traversal of its nodes' values.

func postorder(root *Node) []int {
	ans := make([]int, 0)
	postorderLoop(root, &ans)
	return ans
}

func postorderLoop(node *Node, ans *[]int) {
	if node == nil {
		return
	}
	for _, v := range node.Children {
		postorderLoop(v, ans)
	}
	*ans = append(*ans, node.Val)
}
