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
