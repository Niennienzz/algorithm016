package solution

type stringQueue struct {
	data []string
}

func newStringQueue(data []string) *stringQueue {
	return &stringQueue{
		data: data,
	}
}

func (x stringQueue) len() int {
	return len(x.data)
}

func (x stringQueue) isEmpty() bool {
	return x.len() == 0
}

func (x *stringQueue) push(v string) {
	x.data = append(x.data, v)
}

func (x *stringQueue) pop() string {
	v := x.data[0]
	x.data = x.data[1:]
	return v
}
