package solution

// Given an integer array with no duplicates.
// A maximum tree building on this array is defined as follow:
//  - The root is the maximum number in the array.
//  - The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
//  - The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
// Construct the maximum tree by the given array and output the root node of this tree.
//
// Examples:
//
// Input: [3,2,1,6,0,5]
// Output: return the tree root node representing the following tree:
//
//       6
//     /   \
//    /     \
//   3       5
//    \     /
//     2   0
//       \
//        1

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTreeLoop(nums, 0, len(nums))
}

func constructMaximumBinaryTreeLoop(nums []int, i, j int) *TreeNode {
	if i == j {
		return nil
	}
	maxIndex := maxIndexInRange(nums, i, j)
	node := &TreeNode{Val: nums[maxIndex]}
	node.Left = constructMaximumBinaryTreeLoop(nums, i, maxIndex)
	node.Right = constructMaximumBinaryTreeLoop(nums, maxIndex+1, j)
	return node
}

func maxIndexInRange(nums []int, i, j int) int {
	max := i
	for idx := i; idx < j; idx++ {
		if nums[idx] > nums[max] {
			max = idx
		}
	}
	return max
}
