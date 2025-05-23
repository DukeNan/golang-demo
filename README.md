# Go语言编程知识体系

## 基本数据类型
  - **字节（byte）**
  - **整型（int）**
    - 有符号（signed int）
    - 无符号（unsigned int）
  - **浮点型（float）**
    - 单精度浮点（float32）
    - 双精度浮点（float64）
  - **字符串（string）**
  - **布尔（bool）**

## 基本语法
  - **运算符**
    - 算数运算符
    - 赋值运算符
    - 关系运算符
    - 逻辑运算符
    - 位运算符
    - 特殊运算符
      - 取值运算符（`*`）
      - 通道运算符（`<-`）
    - 运算符优先级
      - 单目运算符
      - 双目运算符
  - **流程控制**
    - 条件判断（`if`）
    - `switch`语句
    - 循环语句
      - `for`
      - 嵌套循环
      - 跳出控制
        - `break`
        - `continue`
        - `goto`
  - **函数**
    - 不定参函数
    - 函数类型
    - 多文件编程

## 复合数据类型
  - **数组（array）**
  - **切片（slice）**
  - **字典（map）**
  - **结构体（struct）**
    - 匿名字段
    - 方法
      - 方法继承
      - 方法重写
  - **指针（pointer）**
  - **接口（interface）**
    - 空接口
    - 类型断言
    - 类型转换

## 并发编程
  - **Channel**
    - 单向 Channel
    - 双向 Channel
    - `select`语句
    - 关闭通道
  - **同步机制**
    - 互斥锁（`sync.Mutex`）
    - 临时对象池（`sync.Pool`）
  - **Go协程（Goroutine）**
  - **底层实现**
    - 调度模型（GMP）
    - 内存管理

## 开发与编译
  - **错误处理**
    - `error`接口
    - `panic`异常
    - `defer`延迟调用
  - **CGO混合编程**
    - `import "C"`
  - **依赖管理**
    - `go mod`
    - `go vendor`
  - **内存分析**
  - **运行模式**
    - 逃逸分析
    - GC机制

## 其他核心概念
  - **面向对象**
  - **包管理**
  - **代码区与数据区**
    - 初始化数据区
    - 代码区