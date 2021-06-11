package main

import (
	"fmt"
	"math"
)

func main() {
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
	c := math.Ceil(1.49)
	fmt.Printf("%T\n", c)
	fmt.Println("-------", c)
	fmt.Printf("%d\n", int(c))
	fmt.Printf("%.1f\n", c)
	fmt.Printf("%.2f\n", math.Copysign(3.2, -1))
	fmt.Println(math.Dim(-2, 4))
	fmt.Println(math.Exp2(2))
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Max(3, 5.6))
	fmt.Println(math.Mod(3, 2))
	fmt.Println(math.Modf(3.14))
	fmt.Println(math.Round(3.14))
	fmt.Println(math.Round(-3.7))
	fmt.Printf("%.1f\n", math.RoundToEven(11.5))
	fmt.Printf("%.1f\n", math.RoundToEven(12.5))
	fmt.Printf("%.1f\n", math.Sqrt(4))
}
