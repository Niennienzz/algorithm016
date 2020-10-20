package solution

// Write a function to delete a node in a singly-linked list.
// You will not be given access to the head of the list.
// Instead you will be given access to the node to be deleted directly.
// It is guaranteed that the node to be deleted is not a tail node in the list.
//
// Examples:
//
// Input: head = [4,5,1,9], node = 5
// Output: [4,1,9]
//
// Input: head = [1,2,3,4], node = 3
// Output: [1,2,4]
//
// Input: head = [-3,5,-99], node = -3
// Output: [5,-99]

func deleteNode(node *ListNode) {
	prev, curr := node, node

	// The node to be deleted is not a tail node in the list.
	curr.Val = curr.Next.Val
	curr = curr.Next

	for curr.Next != nil {
		curr.Val = curr.Next.Val
		prev = prev.Next
		curr = curr.Next
	}
	prev.Next = nil
}
