package solution

// General tree node.
type Node struct {
	Val      int
	Children []*Node
}

type nodeQueue struct {
	data []*Node
}

func newNodeQueue() *nodeQueue {
	return &nodeQueue{
		data: make([]*Node, 0),
	}
}

func (x *nodeQueue) isEmpty() bool {
	return len(x.data) == 0
}

func (x *nodeQueue) size() int {
	return len(x.data)
}

func (x *nodeQueue) push(v *Node) {
	x.data = append(x.data, v)
}

func (x *nodeQueue) pushAll(v []*Node) {
	x.data = append(x.data, v...)
}

func (x *nodeQueue) pop() *Node {
	if x.isEmpty() {
		return nil
	}
	v := x.data[0]
	x.data = x.data[1:]
	return v
}
