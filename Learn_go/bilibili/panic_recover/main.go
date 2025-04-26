package main

import "fmt"

func a() {
	fmt.Println("func a")
}

func b() {
	//recover()必须搭配defer使用
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}
func main() {
	a()
	b()
	c()
}
