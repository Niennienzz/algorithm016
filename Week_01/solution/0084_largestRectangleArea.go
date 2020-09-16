package solution

// Given n non-negative integers representing the histogram's bar height
// where the width of each bar is 1, find the area of largest rectangle in the histogram.
//
// Examples:
//
// Input: [2,1,5,6,2,3]
// Output: 10
//

func largestRectangleAreaUsingStack(heights []int) int {
	var (
		max   = 0
		stack = newIntStack()
	)

	stack.push(-1)

	for k := range heights {
		for stack.top() != -1 && heights[stack.top()] >= heights[k] {
			max = maxInt(max, heights[stack.pop()]*(k-stack.top()-1))
		}
		stack.push(k)
	}

	for stack.top() != -1 {
		max = maxInt(max, heights[stack.pop()]*(len(heights)-stack.top()-1))
	}

	return max
}
