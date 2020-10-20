package solution

// Merge two sorted linked lists and return it as a new sorted list.
// The new list should be made by splicing together the nodes of the first two lists.
//
// Examples:
//
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		prev = &ListNode{0, nil}
		curr = prev
	)

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 == nil {
		curr.Next = l2
	} else {
		curr.Next = l1
	}

	return prev.Next
}
