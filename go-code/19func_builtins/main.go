package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		// 如果程序出现了Panic错误，可以通过recover恢复过来
		if err != nil {
			// fmt.Println(err)s
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("funx C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
