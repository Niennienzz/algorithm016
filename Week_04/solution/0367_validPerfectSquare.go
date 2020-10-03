package solution

// Given a positive integer num, write a function which returns true if num is a perfect square else false.
//
// Examples:
//
// Input: num = 16
// Output: true
//
// Input: num = 16
// Output: true

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	left, right := 2, num/2
	var mid, guess int

	for left <= right {
		mid = left + (right-left)/2
		guess = mid * mid
		if guess == num {
			return true
		}
		if guess > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
