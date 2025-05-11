package main

import (
	"fmt"
	"reflect"
)

type MyInt int64

type person struct {
	name string
	age  int
}

type book struct{ title string }
// 获取类型
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v, kind: %v\n", v.Name(), v.Kind())
}

// 获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))

	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

// 通过反射设置指针变量的值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
	fmt.Println("=============")
	var p *float32
	var m MyInt
	var c rune
	reflectType(p)
	reflectType(m)
	reflectType(c)
	d := person{
		name: "老王",
		age:  20,
	}
	e := book{title: "颈椎康复指南"}
	reflectType(d)
	reflectType(e)
	var s []string
	reflectType(s)
	fmt.Println("-------ValueOf-----------")
	var f1 float32 = 3.14
	reflectValue(f1)
	var i int64 = 100
	fmt.Printf("%T\n", i)
	reflectValue(i)
	h := reflect.ValueOf(10)
	fmt.Printf("type h :%T\n", h) // type c :reflect.Value
	fmt.Println("=============通过反射设置变量的值===========")
	var q int64 = 100
	// reflectSetValue1(q) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&q)
	fmt.Println(q)
}
