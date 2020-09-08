package solution

// Description:
// Given an array nums, write a function to move all zeros to the end of it
// while maintaining the relative order of the non-zero elements.
//
// Examples:
//
// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.

// Analysis:
// The 2 requirements of the question are:
//  - Move all the 0's to the end of array.
//  - All the non-zero elements must retain their original order.
// It is good to realize here that both the requirements are mutually exclusive.
// You can solve the individual sub-problems and then combine them for the final solution.

// Naive solution: count number of zeros, fill non-zero values, and then fill zeros.
// Space Complexity: O(n). Since we are creating the "ans" array to store results.
// Time Complexity: O(n). However, the total number of operations are sub-optimal.
func moveZeroes_extraSpace(nums []int) {
	numOfZeros := 0

	for _, v := range nums {
		if v == 0 {
			numOfZeros += 1
		}
	}

	ans := make([]int, 0)
	for _, v := range nums {
		if v != 0 {
			ans = append(ans, v)
		}
	}

	for numOfZeros != 0 {
		ans = append(ans, 0)
		numOfZeros -= 1
	}

	for k := range nums {
		nums[k] = ans[k]
	}
}

// Rephrase the problem: bring all the non-zero elements to the front of array keeping their relative order same.
// Space Complexity: O(1). Only constant space is used.
// Time Complexity:  O(n). However, the total number of operations are still sub-optimal.
// The total operations (array writes) that code does is n (i.e. total number of elements).
// This is a two-pointer solution, where the slow pointer starts only when the fast pointer finishes the job.
func moveZeroes_extraOperations(nums []int) {
	lastNonZeroFoundAt := 0

	for k := range nums {
		if nums[k] != 0 {
			nums[lastNonZeroFoundAt] = nums[k]
			lastNonZeroFoundAt += 1
		}
	}

	for k := lastNonZeroFoundAt; k < len(nums); k++ {
		nums[k] = 0
	}
}

// Optimal solution: when we encounter a non-zero element, we need to swap elements pointed by current and slow pointer,
// then advance both pointers. If it is a zero element, we just advance current pointer.
// Space Complexity: O(1). Only constant space is used.
// Time Complexity: O(n). However, the total number of operations are optimal.
// The total operations is the number of non-0 elements.
// This gives us a much better best-case (when most of the elements are zeros) complexity than last solution.
// The worst-case (when all elements are non-zero elements) complexity for both the algorithms is same.
func moveZeroes_optimal(nums []int) {
	for lastNonZeroFoundAt, k := 0, 0; k < len(nums); k++ {
		if nums[k] != 0 {
			nums[lastNonZeroFoundAt], nums[k] = nums[k], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt += 1
		}
	}
}
