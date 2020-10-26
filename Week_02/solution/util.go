package solution

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minInt3(a, b, c int) int {
	return minInt(a, minInt(b, c))
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
