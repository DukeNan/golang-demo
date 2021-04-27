package main

import "fmt"

func main() {
	// age := 15
	// if age > 18 {
	// 	fmt.Println("已成年")
	// } else {
	// 	fmt.Println("未成年🔞")
	// }

	// if age > 35 {
	// 	fmt.Println("中年")
	// } else if age > 18 {
	// 	fmt.Println("青年")
	// } else {
	// 	fmt.Println("小朋友")
	// }

	// 作用域
	// 此时age只在if语句中生效
	if age := 15; age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("小朋友")
	}

}
