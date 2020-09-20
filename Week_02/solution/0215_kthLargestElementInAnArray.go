package solution

import (
	"container/heap"
	"sort"
)

// Find the kth largest element in an unsorted array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
// Examples:
//
// Input: [3,2,1,5,6,4] and k = 2
// Output: 5
//
// Input: [3,2,3,1,2,4,5,5,6] and k = 4
// Output: 4

func findKthLargestUsingSort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// Use a min-heap, but keep popping if its size is greater than k.
// The resulting heap is a min-heap with only largest elements in ASC order.
func findKthLargestUsingHeap(nums []int, k int) int {
	hp := &intMinHeap{}
	heap.Init(hp)

	for _, v := range nums {
		heap.Push(hp, v)
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}

	return (*hp)[0]
}
