package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 一组基础示例，展示 %v 是默认格式，在此例中整数为十进制显示，也可以用 %d 明确指定；输出与 Println 生成的内容一致。
	integer := 23
	// 下面每个都会打印 "23"（不带引号）。
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	// 特殊的 %T 占位符显示的是变量的类型而不是值。
	fmt.Printf("%T %T\n", integer, &integer)
	// 结果：int *int

	// Println(x) 等价于 Printf("%v\n", x)，所以下面我们只用 Printf。
	// 每个例子演示如何格式化某种类型的值，比如整数或字符串。我们每个格式字符串都以 %v 开头，展示默认输出，然后跟一个或多个自定义格式。

	// 布尔值用 %v 或 %t 打印为 "true" 或 "false"。
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// 结果：true true

	// 整数用 %v 和 %d 打印为十进制，
	// 用 %x 打印为十六进制，用 %o 打印为八进制，用 %b 打印为二进制。
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// 结果：42 42 2a 52 101010

	// 浮点数有多种格式：%v 和 %g 打印紧凑表示，
	// %f 打印小数点，%e 用科学计数法。
	// %6.2f 这里展示了如何设置宽度和精度来控制浮点数的显示。
	// 这里 6 是总宽度（注意输出中的空格），2 是小数位数。
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// 结果：3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00

	// 复数格式为括号包裹的浮点数对，虚部后有 'i'。
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// 结果：(110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)

	// rune 是整数，但用 %c 打印时显示对应的 Unicode 字符。
	// %q 显示为带引号的字符，%U 以十六进制 Unicode 码点显示，%#U 既显示码点又显示可打印字符。
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// 结果：128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

	// 字符串用 %v 和 %s 原样输出，%q 加引号，%#q 用反引号。
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// 结果：foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

	// map 用 %v 格式化时显示键值对的默认格式。
	// %#v 以 Go 源码格式显示 map。
	// map 打印顺序一致，按键排序。
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// 结果：map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	// 结构体用 %v 显示字段值的默认格式。
	// %+v 显示字段名，%#v 以 Go 源码格式显示结构体。
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// 结果：{Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	// 指针的默认格式显示底层值并带有 &。
	// %p 以十六进制显示指针值。这里用类型为 nil 的指针作为 %p 的参数，因为非 nil 指针的值每次运行都不同；如需查看效果可自行取消注释。
	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	// 结果：&{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// 结果：&{Kim 22} 0x010203 // 见上面注释。

	// 数组和切片会对每个元素应用格式化。
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// 结果：[Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// 结果：[Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}

	// 字节切片比较特殊。用整数格式（如 %d）打印元素。
	// %s 和 %q 把切片当作字符串。
	// %x 有特殊形式，带空格标志会在字节间加空格。
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// 结果：[97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

	// 实现了 Stringer 接口的类型打印效果与字符串相同。
	// 因为 Stringer 返回字符串，可以用字符串专用的格式化符号如 %q。
	now := time.Unix(123456789, 0).UTC() // time.Time implements fmt.Stringer.
	fmt.Printf("%v %q\n", now, now)
	// 结果：1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"

}

/*
Output:

23
23
23
int *int
true true
42 42 2a 52 101010
3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00
(110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
128512 128512 😀 '😀' U+1F600 U+1F600 '😀'
foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}
{Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
&{Kim 22} 0x0
[Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]
[Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}
[97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"
*/
