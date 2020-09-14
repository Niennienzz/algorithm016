package solution

// Given an array nums, there is a sliding window of size k which is moving from the
// very left of the array to the very right. You can only see the k numbers in the window.
// Each time the sliding window moves right by one position. Return the max sliding window.
//
// Follow up: Could you solve it in linear time?
//
// Examples:
//
// Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// -----------------------------------
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

func maxSlidingWindowNaive(nums []int, k int) []int {
	var (
		n   = len(nums)
		ans = make([]int, n-k+1)
	)

	if n*k == 0 {
		return nil
	}

	for i := 0; i < n-k+1; i++ {
		max := -2147483648
		for j := i; j < k+i; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		ans[i] = max
	}

	return ans
}

func maxSlidingWindowDeque(nums []int, k int) []int {
	var (
		n     = len(nums)
		deque = NewMyCircularDeque(k)
	)

	if n*k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	maxIndex := 0
	for i := 0; i < k; i++ {
		cleanDeque(i, k, nums, &deque)
		deque.InsertLast(i)
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	ans := make([]int, n-k+1)
	ans[0] = nums[maxIndex]

	for i := k; i < n; i++ {
		cleanDeque(i, k, nums, &deque)
		deque.InsertLast(i)
		ans[i-k+1] = nums[deque.GetFront()]
	}
	return ans
}

func cleanDeque(i, k int, nums []int, deque *MyCircularDeque) {
	if !deque.IsEmpty() && deque.GetFront() == i-k {
		deque.DeleteFront()
	}
	for !deque.IsEmpty() && nums[i] > nums[deque.GetRear()] {
		deque.DeleteLast()
	}
}
