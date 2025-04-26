package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	// 不知道别人调用这个函数时会传入什么类型的变量
	// 1.方法1：通过类型断言
	// 2.方法2：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
	fmt.Printf("%T\n", obj) // *reflect.type
}

func main() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 10
	reflectType(b)
}
