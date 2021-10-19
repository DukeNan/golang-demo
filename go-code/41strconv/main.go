package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("============Atoi======================")
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int ")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1)
	}
	fmt.Println("=============Itoa=====================")
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)
	fmt.Println("=============Parse=====================")
	b, err := strconv.ParseBool("true")
	if err != nil {
		return
	}
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	i, err := strconv.ParseInt("-2", 10, 64)
	if err != nil {
		return
	}
	u, err := strconv.ParseUint("2", 10, 64)
	if err != nil {
		return
	}
	fmt.Printf("b=%v, f=%v, i=%v, u=%v\n", b, f, i, u)
	fmt.Println("=============Format=====================")
	s11 := strconv.FormatBool(true)
	s12 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s13 := strconv.FormatInt(-2, 16)
	s14 := strconv.FormatUint(2, 16)
	fmt.Printf("s11=%v, s12=%v, s13=%v, s14=%v\n", s11, s12, s13, s14)

}
