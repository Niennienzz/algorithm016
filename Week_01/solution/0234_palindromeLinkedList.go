package solution

// Given a singly linked list, determine if it is a palindrome.
// Could you do it in O(n) time and O(1) space?
//
// Examples:
//
// Input: 1->2
// Output: false
//
// Input: 1->2->2->1
// Output: true

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 注意和题目 876 的区别.
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var (
		ans         = true
		firHalfTail = slow
		secHalfHead = reverseListIterative(slow.Next)
	)

	p1, p2 := head, secHalfHead
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			ans = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	firHalfTail.Next = reverseListIterative(secHalfHead)
	return ans
}
