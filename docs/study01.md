## 变量和常量

Go 语言中的变量必须先声明再使用

### 关键字

Go语言关键字

```go
break      default      func     interface    select
case       defer        go       map          struct
chan       else         goto     package      switch
const      fallthrough  if       range        type
continue   for          import   return       var
```

 此外，还有37个保留字

```go
Constants:    true  false     iota    nil

Types:   int      int8         int16        int32      int64
				 uint     uint8        uint16       uint32     uint64    uintptr
				 float32  float64      complex128   complex64
				 bool     byte         rune         string     error

Functions:	 make  len cap  new  append   copy    close   delete
						 complex    real      imag 
						 panic      recover
```

### 声明变量

`var s1 string ` 声明一个保存字符串类型数据的s1变量

```go
package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "laowang"
	age = 18
	isOk = true
	// Go语言中变量声明必须使用， 不使用编译不过去
	fmt.Print(isOk)              // 在终端输出要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n", name) //  格式化输出
	fmt.Println(age)             // 打印之后换行
	fmt.Println("hello world")

	//  声明变量同时赋值
	var s1 string = "lisi"
	fmt.Println(s1)

	// 类型推导（根据值判断该变量的类型）
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明,  只能在函数里面声明
	s3 := "hahaha"
	fmt.Println(s3)

	// S1 := "10"  //   同一个作用域不能重复声明同名变量

	//  匿名变量是一个特殊的变量：_
}

```

#### iota

`iota`是够语言的常量计数器，只能在常量的表达式中使用。

`iota`在`const`关键字出现时被重置为0，`const`中每增加一行常量声明将使`iota`计数一次（`iota`可理解为`const`语句块中的索引）。使用`iota`能简化定义，在定义枚举时很有用。

举个例子：

```go
const (
  n1 = iota   //  0
  n2          // 1
  n3          // 2
  n4          // 3
)
```

使用`_`跳过某些值

```go
const (
		n1 = iota //0
		n2        //1
		_
		n4        //3
	)
```

`iota`声明中间插队

```go
const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
const n5 = iota //0
```

定义数量级 （这里的`<<`表示左移操作，`1<<10`表示将1的二进制表示向左移10位，也就是由`1`变成了`10000000000`，也就是十进制的1024。同理`2<<2`表示将2的二进制表示向左移2位，也就是由`10`变成了`1000`，也就是十进制的8。）

```go
const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
```

多个`iota`定义在一行

```go
const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
```

