package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for循环的初始语句和结束语句都可以省略
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// s := "hello"
	// fmt.Println(len(s))
	// for i, v := range s {
	// 	// fmt.Println(i, v)
	// 	fmt.Printf("%d %c\n", i, v)
	// }

	// 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%-3d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
