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

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n*k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = maxInt(left[i-1], nums[i])
		}

		j := n - i - 1
		if (j+1)%k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = maxInt(right[j+1], nums[j])
		}
	}

	ans := make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		ans[i] = maxInt(left[i+k-1], right[i])
	}
	return ans
}

func maxSlidingWindowDeque(nums []int, k int) []int {
	ans := make([]int, 0)
	idx := make([]int, 0)
	for i, v := range nums {
		for len(idx) != 0 && v >= nums[idx[len(idx)-1]] {
			idx = idx[:len(idx)-1]
		}
		idx = append(idx, i)
		if idx[0] == i-k {
			idx = idx[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[idx[0]])
		}
	}
	return ans
}
