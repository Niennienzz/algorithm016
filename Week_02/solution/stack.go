package solution

// Integer.
type intStack struct {
	data []int
}

func newIntStack() *intStack {
	return &intStack{
		data: make([]int, 0),
	}
}

func (x *intStack) push(v int) {
	x.data = append(x.data, v)
}

func (x *intStack) pop() int {
	v := x.data[len(x.data)-1]
	x.data = x.data[:len(x.data)-1]
	return v
}

func (x *intStack) top() int {
	return x.data[len(x.data)-1]
}

func (x *intStack) isEmpty() bool {
	return len(x.data) == 0
}

// TreeNode.
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
