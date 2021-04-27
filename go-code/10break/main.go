package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	var n = 3
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效数字")
	}

	if age := 20; age > 18 && age < 60 {
		fmt.Println("人到中年不得已，保温杯里泡枸杞！")
	}

	var m = 10
	fmt.Printf("%T\n", m)

	var x int
	x = 10
	x += 1
	fmt.Println(x)
}
