package solution

// 1 -> 2 -> 3 -> nil
func reverseListIterative(head *ListNode) *ListNode {
	var (
		prev *ListNode
		curr = head
	)
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
