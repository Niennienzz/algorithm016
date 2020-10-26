package solution

// Binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack.
type treeNodeStack struct {
	data []*TreeNode
}

func newTreeNodeStack() *treeNodeStack {
	return &treeNodeStack{
		data: make([]*TreeNode, 0),
	}
}

func (x *treeNodeStack) push(v *TreeNode) {
	x.data = append(x.data, v)
}

func (x *treeNodeStack) pop() *TreeNode {
	v := x.data[len(x.data)-1]
	x.data = x.data[:len(x.data)-1]
	return v
}

func (x *treeNodeStack) top() *TreeNode {
	return x.data[len(x.data)-1]
}

func (x *treeNodeStack) isEmpty() bool {
	return len(x.data) == 0
}

// Queue.
type treeNodeQueue struct {
	data []*TreeNode
}

func newTreeNodeQueue() *treeNodeQueue {
	return &treeNodeQueue{
		data: make([]*TreeNode, 0),
	}
}

func (x treeNodeQueue) size() int {
	return len(x.data)
}

func (x treeNodeQueue) isEmpty() bool {
	return x.size() == 0
}

func (x *treeNodeQueue) pop() *TreeNode {
	v := x.data[0]
	x.data = x.data[1:]
	return v
}

func (x *treeNodeQueue) push(v *TreeNode) {
	x.data = append(x.data, v)
}
