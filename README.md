# My LeetCode Go Project

这个仓库包含了我的 LeetCode 题目解答，使用 Go 语言实现。

## 目录结构

- `go.mod`: Go 模块文件。
- `README.md`: 本说明文件。
- `problems/`: 存放所有 LeetCode 题解。
  - `xxxx_problem_name/`: 每个题目的独立文件夹，`xxxx` 是题目编号。
    - `problem_name.go`: 该题目的 Go 解答代码。
    - `problem_name_test.go`: 对应的单元测试。
- `pkg/`: 存放可复用的公共包。
  - `utils/`: 示例通用工具包。
    - `utils.go`
    - `utils_test.go`
  - `datastructures/`: 存放常用的数据结构实现。
    - `linkedlist/`
    - `tree/`
    - `(更多...)`

## 如何使用

### 运行测试
由于没有了统一的 `main` 入口，主要通过测试来验证代码。
在题目对应的文件夹内，例如 `problems/0001_two_sum/`，运行：
```bash
go test
```