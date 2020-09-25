学习笔记

## 递归

- 递归的代码模板

```python
def recursion(level, param1, param2, ...):
    # 1 - Recursion terminator.
    if level > MAX_LEVEL:
        process_result()
        return

    # 2 - Process logic in the current level.
    process(level, data...)

    # 3 - Drill down.
    self.recursion(level+1, p1, p2...)

    # 4 - Reverse the current level status if need.
```

- 要点:
  - 不要进行人肉递归. (熟练之后, 不要过多地去画状态树, 不要试图去想象每一个步骤)
  - 找到最近最简单的方法, 将其拆解成可重复解决的子问题.
  - 数学归纳法的思维.

## 分治

- 分治的代码模板

```python
def divide_conquer(problem, param1, param2, ...):
    # 1 - Recursion terminator.
    if problem is None: 
        print_result
        return 

    # 2 - Prepare data.
    data = prepare_data(problem)
    subproblems = split_problem(problem, data) 

    # 3 - Conquer subproblems. (与简单递归相比, 3/4 要把子问题的结果合并起来.)
    subresult1 = self.divide_conquer(subproblems[0], p1, ...) 
    subresult2 = self.divide_conquer(subproblems[1], p1, ...) 
    subresult3 = self.divide_conquer(subproblems[2], p1, ...) 
    ...

    # 4 - Process and generate the final result. (与简单递归相比, 3/4 要把子问题的结果合并起来.)
    result = process_result(subresult1, subresult2, subresult3, …)

    # 5 - Revert the current level states.
```

## 回溯

- 在每一层去试错, 分步地去解决一个问题.
- 错误的话, 取消上一步甚至上几步的计算.
- 通常用递归方法来实现, 在反复上述步骤后:
  - 找到一个可能存在的正确答案.
  - 在尝试了所有可能的分步方法后宣告该问题没有答案.
- 最坏情况下, 回溯法会导致一次复杂度为指数时间的计算.
- 经典问题: **N皇后**, **数独**.
