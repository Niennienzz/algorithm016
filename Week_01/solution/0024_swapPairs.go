package solution

// Given a linked list, swap every two adjacent nodes and return its head.
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Examples:
//
// Given 1->2->3->4, you should return the list as 2->1->4->3.

func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first, second := head, head.Next
	first.Next = swapPairsRecursive(second.Next)
	second.Next = first

	return second
}

func swapPairsIterative(head *ListNode) *ListNode {
	var (
		preHead = new(ListNode)
		prev    = preHead
	)
	preHead.Next = head

	for head != nil && head.Next != nil {
		// Nodes to be swapped.
		first, second := head, head.Next

		// Swap.
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		// Prepare for the next iteration.
		prev = first
		head = first.Next
	}

	return preHead.Next
}
