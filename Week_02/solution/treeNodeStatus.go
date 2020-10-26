package solution

type treeNodeStatus bool

const (
	statusNewNode treeNodeStatus = false
	statusVisited treeNodeStatus = true
)

type treeNodeWithStatus struct {
	*TreeNode
	status treeNodeStatus
}

func newTreeNodeWithStatus(treeNode *TreeNode) *treeNodeWithStatus {
	if treeNode == nil {
		return nil
	}
	return &treeNodeWithStatus{
		TreeNode: treeNode,
		status:   statusNewNode,
	}
}

type treeNodeWithStatusStack struct {
	data []*treeNodeWithStatus
}

func newTreeNodeWithStatusStack() *treeNodeWithStatusStack {
	return &treeNodeWithStatusStack{
		data: make([]*treeNodeWithStatus, 0),
	}
}

func (x *treeNodeWithStatusStack) push(v *treeNodeWithStatus) {
	x.data = append(x.data, v)
}

func (x *treeNodeWithStatusStack) pop() *treeNodeWithStatus {
	v := x.data[len(x.data)-1]
	x.data = x.data[:len(x.data)-1]
	return v
}

func (x *treeNodeWithStatusStack) top() *treeNodeWithStatus {
	return x.data[len(x.data)-1]
}

func (x *treeNodeWithStatusStack) isEmpty() bool {
	return len(x.data) == 0
}

func (root *TreeNode) InorderTraversal() []int {
	var (
		ans   []int
		stack = newTreeNodeWithStatusStack()
	)

	stack.push(newTreeNodeWithStatus(root))

	for !stack.isEmpty() {
		node := stack.pop()
		if node == nil {
			continue
		}
		if node.status == statusNewNode {
			node.status = statusVisited
			// 不同顺序的遍历, 只需调整以下三个语句的顺序即可.
			stack.push(newTreeNodeWithStatus(node.Right))
			stack.push(node)
			stack.push(newTreeNodeWithStatus(node.Left))
		} else {
			ans = append(ans, node.Val)
		}
	}

	return ans
}

func (root *TreeNode) PreorderTraversal() []int {
	var (
		ans   []int
		stack = newTreeNodeWithStatusStack()
	)

	stack.push(newTreeNodeWithStatus(root))

	for !stack.isEmpty() {
		node := stack.pop()
		if node == nil {
			continue
		}
		if node.status == statusNewNode {
			node.status = statusVisited
			// 不同顺序的遍历, 只需调整以下三个语句的顺序即可.
			stack.push(newTreeNodeWithStatus(node.Right))
			stack.push(newTreeNodeWithStatus(node.Left))
			stack.push(node)
		} else {
			ans = append(ans, node.Val)
		}
	}

	return ans
}

func (root *TreeNode) PostorderTraversal() []int {
	var (
		ans   []int
		stack = newTreeNodeWithStatusStack()
	)

	stack.push(newTreeNodeWithStatus(root))

	for !stack.isEmpty() {
		node := stack.pop()
		if node == nil {
			continue
		}
		if node.status == statusNewNode {
			node.status = statusVisited
			// 不同顺序的遍历, 只需调整以下三个语句的顺序即可.
			stack.push(node)
			stack.push(newTreeNodeWithStatus(node.Right))
			stack.push(newTreeNodeWithStatus(node.Left))
		} else {
			ans = append(ans, node.Val)
		}
	}

	return ans
}
