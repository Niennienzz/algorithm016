# 学习笔记

## 搜索

- 数据结构如果没有特点 (普通的树/图) 暴力搜索: 其实就是要把所有节点访问一次, 且仅访问一次.
- 对于节点的顺序不同:
  - 深度优先
  - 广度优先

### 深度优先 (DFS) 模板

- 递归

```python
visited = set()

def DFS(node, visited):
    # Terminator.
    if node in visited:
        # Already visited!
        return

    visited.add(node)

    # Process current node here.
    ...
    for next_node in node.children():
        if next_node not in visited:
            dfs(next_node, visited)
```

- 非递归 (手动维护栈)

```python
def DFS(self, tree):
    if tree.root is None:
        return []

    visited, stack = [], [tree.root]

    while stack:
        node = stack.pop()
        visited.add(node)

        process (node)
        nodes = generate_related_nodes(node)
        stack.push(nodes)

    # Other processing work.
    ...
```

### 广度优先 (BFS) 模板

- 手动维护队列

```python
def BFS(graph, start, end):
    visited = set()
    queue = []
    queue.append([start])
    while queue:
        node = queue.pop()
        visited.add(node)
        process(node)
        nodes = generate_related_nodes(node)
        queue.push(nodes)

    # Other processing work.
    ...
```

## 贪心算法

- 在每一步选择中, 都选择当前状态下的最好或最优, 从而希望导致结果是全局的最好或最优结果.
- 对比动态规划:
  - 贪心算法不能退回.
  - 动态规划则会保存以前的运算结果, 并根据以前的结果对当前进行选择, 有退回功能.
- 贪心算法可以解决一些最优化问题, 如: **求图中的最小生成树**, **求霍夫曼编码**等.
- 然而对于工程和生活中的问题, 贪心算法一般不能得到我们所要求的答案.
- 一旦一个问题可以通过贪心算法来解决, 那么贪心算法一般是解决这个问题的最好办法.
- 由于贪心算法的高效性以及其所求得的答案比较接近最优结果:
  - 贪心算法也可以用作辅助算法.
  - 或者直接解决一些要求结果不是特别精确的问题.
- 贪心算法难在: **如何证明使用贪心算法能得到全局最优解**.
  - 例如找零钱问题中, 假设零钱面值为: \[10, 5, 1], 我们需要找最少的零钱个数给对方.
  - 因为零钱的三个面值它们**成倍数关系**, 两个 5 肯定比不上一个 10, 以此类推, 所以可以用贪心算法.

## 二分查找

- 先决条件:
  - 有序: 单调递增或递减.
  - 存在上下界.
  - 能够通过索引访问 (单链表 vs 跳表)
  
### 二分查找模板

```python
left, right = 0, len(array) - 1
while left <= right:
    mid = (left + right) / 2
    if array[mid] == target:
        break or return result
    elif array[mid] < target:
        left = mid + 1
    else:
        right = mid - 1
```
