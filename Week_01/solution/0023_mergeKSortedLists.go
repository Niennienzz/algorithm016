package solution

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
//
// Examples:
//
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
//
// Input: lists = []
// Output: []
//
// Input: lists = [[]]
// Output: []

// Use mergeTwoLists as a helper.
// Merge list one-by-one.
// Time complexity: O(kN)
// Space complexity: O(1)
func mergeKListsUsingMergeTwoLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for _, list := range lists {
		ans = mergeTwoLists(ans, list)
	}
	return ans
}

// Use mergeTwoLists as a helper.
// However, merge a pair of lists each time by divide and conquer.
// Time complexity: O(N * log k) where k is the number of linked lists.
// Space complexity: O(1)
func mergeKListsUsingDivideAndConquer(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	step := 1
	for step < n {
		for i := 0; i < n-step; i += (step * 2) {
			lists[i] = mergeTwoLists(lists[i], lists[i+step])
		}
		step *= 2
	}
	return lists[0]
}
