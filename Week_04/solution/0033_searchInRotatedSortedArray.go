package solution

// You are given an integer array nums sorted in ascending order, and an integer target.
// Suppose that nums is rotated at some pivot unknown to you beforehand ([0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// If target is found in the array return its index, otherwise, return -1.
//
// Examples:
//
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
//
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
//
// Input: nums = [1], target = 0
// Output: -1

// 二分查找, 注意边界条件.
func search(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
		mid   int
	)
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
