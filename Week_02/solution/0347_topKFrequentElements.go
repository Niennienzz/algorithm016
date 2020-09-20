package solution

import (
	"container/heap"
)

// Given a non-empty array of integers, return the k most frequent elements.
//
// Examples:
//
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
//
// Input: nums = [1], k = 1
// Output: [1]

// Use a min-heap, but keep popping if its size is greater than k.
// The resulting heap is a min-heap with only largest elements in ASC order.
func topKFrequent(nums []int, topK int) []int {
	// Count using a map.
	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}

	// Push count-value pairs into a heap, keep top K.
	hp := &vcPairMinHeap{}
	heap.Init(hp)
	for k, v := range m {
		p := vcPair{value: k, count: v}
		heap.Push(hp, p)
		if hp.Len() > topK {
			heap.Pop(hp)
		}
	}

	// Create answer from the heap reversely.
	ans := make([]int, topK)
	for k := topK - 1; k >= 0; k-- {
		ans[k] = heap.Pop(hp).(vcPair).value
	}
	return ans
}

type vcPair struct {
	value int
	count int
}

type vcPairMinHeap []vcPair

func (h vcPairMinHeap) Len() int           { return len(h) }
func (h vcPairMinHeap) Less(i, j int) bool { return h[i].count < h[j].count } // Order is based on count.
func (h vcPairMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *vcPairMinHeap) Push(x interface{}) {
	*h = append(*h, x.(vcPair))
}

func (h *vcPairMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
