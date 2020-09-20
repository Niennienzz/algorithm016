package solution

type intMaxHeap []int

func (h intMaxHeap) Len() int           { return len(h) }
func (h intMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
