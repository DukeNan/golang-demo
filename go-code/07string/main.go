package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/Users/dukenan/go/src/github.com/dukenan/golang-demo/go-code"
	// 获取字符串类型
	fmt.Printf("%T\n", s)
	// 获取字符串值
	fmt.Printf("%v\n", s)
	s1 := `
	 abd
	 edf
	 ghi
	`
	fmt.Println(s1)
	// 	字符串相关操作
	s2 := "kshf/hskdfk/akfd"
	fmt.Println(len(s2))
	// 字符串拼接
	name := "laowang"
	age := "19"

	fmt.Println(name + age)
	fmt.Printf("%s%s\n", name, age)
	s3 := fmt.Sprintf("%s%s", name, age)
	fmt.Println(s3)
	// 字符串分割
	ret := strings.Split(s2, "/")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(name, "lao"))
	// 前缀/后缀
	fmt.Println(strings.HasPrefix(name, "lao"))
	fmt.Println(strings.HasSuffix(name, "aaa"))
	// 字串出现的位置
	s4 := "abcdefghcj"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "c"))
	// 拼接
	fmt.Println(strings.Join(ret, "+"))
	fmt.Println("==================================================")

	// len() 求的是byte字节的数量
	s10 := "hello world深圳"
	n := len(s10)
	fmt.Println(n)
	// for i := 0; i < len(s10); i++ {
	// 	// fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s10[i])
	// }

	// for _, c:= range s10 {  // 从字符串中拿出具体的字符
	// 	fmt.Printf("%c\n", c)
	// }

	// 字符串修改
	s12 := "白萝卜"             // [白 萝 卜]
	s13 := []rune(s12)       // 把字符串强制转换成一个rune切片
	s13[0] = '红'             // 切片里面是字符 rune(int32)
	fmt.Println(string(s13)) // 把rune切片强制转换成字符串
	// 类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)

}
