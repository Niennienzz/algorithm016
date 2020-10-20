package solution

// Given a non-empty, singly linked list with head node head, return a middle node of linked list.
// If there are two middle nodes, return the second middle node.
//
// Examples:
//
// Input: [1,2,3,4,5]
// Output: Node 3 from this list.
//
// Input: [1,2,3,4,5,6]
// Output: Node 4 from this list.

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
