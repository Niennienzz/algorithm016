package solution

import (
	"container/heap"
)

// Write a program to find the n-th ugly number.
// Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
//  - 1 is typically treated as an ugly number.
//  - n does not exceed 1690.
//
// Examples:
//
// Input: n = 10
// Output: 12
// Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.

func nthUglyNumberUsingHeap(n int) int {
	var (
		seen = make(map[int]struct{})
		hp   = &intMinHeap{}
		nums = [1690]int{}
	)

	heap.Init(hp)
	hp.Push(1)
	seen[1] = struct{}{}

	currUgly, newUgly := 0, 0
	primes := []int{2, 3, 5}

	for i := 0; i < 1690; i++ {
		currUgly = heap.Pop(hp).(int)
		nums[i] = currUgly
		for _, j := range primes {
			newUgly = currUgly * j
			if _, ok := seen[newUgly]; !ok {
				seen[newUgly] = struct{}{}
				heap.Push(hp, newUgly)
			}
		}
	}

	return nums[n-1]
}
