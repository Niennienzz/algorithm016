package solution

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
// Examples:
//
// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
//  - Rotate 1 steps to the right: 5->1->2->3->4->NULL
//  - Rotate 2 steps to the right: 4->5->1->2->3->NULL
//
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
//  - Rotate 1 steps to the right: 2->0->1->NULL
//  - Rotate 2 steps to the right: 1->2->0->NULL
//  - Rotate 3 steps to the right: 0->1->2->NULL
//  - Rotate 4 steps to the right: 2->0->1->NULL

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Find how long the list is.
	// Close the list into a ring.
	oldTail := head
	var n int
	for n = 1; oldTail.Next != nil; n++ {
		oldTail = oldTail.Next
	}
	oldTail.Next = head

	// New tail is at (n - k % n - 1).
	// New head is at (n - k % n).
	newTail := head
	for i := 0; i < n-k%n-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next

	// Break the ring.
	newTail.Next = nil
	return newHead
}
