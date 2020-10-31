package solution

type DisjointSet struct {
	size int   // 当前集合的个数, 小于或等于元素个数.
	reps []int // 并查集储存元素的实际结构.
}

// 并查集初始化: 每个元素指向自己, 自己是自己的代表, 同时自己也是一个独立集合.
func NewDisjointSet(n int) *DisjointSet {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}
	return &DisjointSet{
		size: n,
		reps: r,
	}
}

// 并查集查找: 如果值的代表不是自己, 则不断寻找上一级代表, 同时压缩路径.
//
// 路径压缩前: 1 <-- 2 <-- 3 <-- 4
//
// 路径压缩后: 1 <-- 2
//          / \
//         /   \
//        3     4
func (x DisjointSet) Find(n int) int {
	// 边界条件, 注意这里用 Cap() 而不是 size, 因为 size 的值是随着 Union() 而变化的.
	if n < 0 || n > x.Cap() {
		panic("target value out of range")
	}

	// 如图 n 的代表不是自己, 则寻找上一级代表.
	// 同时压缩路径: n 的上一级代表, 更新为上上级.
	for n != x.reps[n] {
		x.reps[n] = x.reps[x.reps[n]]
		n = x.reps[n]
	}

	return n
}

// 并查集合并: 找到两个值的最上级代表, 如果不相等, 则使其中一个成为另一个的上级.
func (x *DisjointSet) Union(p, q int) {
	rootP := x.Find(p)
	rootQ := x.Find(q)

	if rootP == rootQ {
		return
	}

	x.reps[rootP] = rootQ
	x.size -= 1
}

func (x DisjointSet) Size() int {
	return x.size
}

func (x DisjointSet) Cap() int {
	return len(x.reps)
}
