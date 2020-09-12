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
		before = &ListNode{0, nil}
		cursor = before
	)

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			l2 = l2.Next
		}
		cursor = cursor.Next
	}

	if l1 == nil {
		cursor.Next = l2
	} else {
		cursor.Next = l1
	}

	return before.Next
}
