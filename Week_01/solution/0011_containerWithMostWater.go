package solution

// Basically list all possible areas.
func maxAreaBruteForce(height []int) int {
	max := 0
	for i := range height {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * minInt(height[i], height[j])
			max = maxInt(max, area)
		}
	}
	return max
}

// Use the Squeeze theorem with a two-pointer method.
func maxArea(height []int) int {
	max, i, j := 0, 0, len(height)-1
	var w, h, area int
	for i != j {
		w = j - i
		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}
		area = w * h
		if area > max {
			max = area
		}
	}
	return max
}
