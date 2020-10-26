package solution

// Given a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.
//
// Examples:
//
// Input: [1,2,3,null,5,null,4]
// Output: [1,3,4]
// Explanation:
//     1       <---
//   /   \
//  2     3    <---
//   \     \
//    5     4  <---
//
// Input: [1,2]
// Output: [1,2]
// Explanation:
//     1  <---
//   /
//  2     <---

// Use a BFS, append the last node value into the answer.
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		ans []int
		que = newTreeNodeQueue()
	)

	que.push(root)

	for !que.isEmpty() {
		size := que.size()
		for i := 0; i < size; i++ {
			node := que.pop()
			if i == size-1 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				que.push(node.Left)
			}
			if node.Right != nil {
				que.push(node.Right)
			}
		}
	}

	return ans
}
