package main

import "fmt"

func main() {
	var a [3]int
	var b [4]int
	a = [3]int{1, 2, 9}
	b = [4]int{1, 3, 5}
	fmt.Println(a, b)
	// 初始化数组时可以使用初始化列表来设置数组元素的值。
	var testArray [3]int
	var numArray = [3]int{1, 2}
	var cityArray = [3]string{"北京", "", ""}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)
	// 一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：
	var testArray1 [3]int
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "广州"}
	fmt.Println(testArray1, numArray1, cityArray1)
	// 我们还可以使用指定索引值的方式来初始化数组
	n := [...]int{0: 100, 5: 110}
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	// =========遍历数组==================
	var m = [...]string{"东莞", "惠州", "中山"}
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
	for index, value := range m {
		fmt.Println(index, value)
	}
	// 多维数组
	t := [3][2]string{
		{"河南", "郑州"},
		{"江苏", "苏州"},
		{"广东", "深圳"},
	}
	// fmt.Println(t)
	// fmt.Println(t[2][1])
	// 遍历二维数组
	for _, v1 := range t {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	// 多维数组只有第一层可以使用...来让编译器推导数组长度
	p := [...][2]string{
		{"河南", "郑州"},
		{"江苏", "苏州"},
		{"广东", "深圳"},
	}
	fmt.Println(p)
}
