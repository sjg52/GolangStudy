package main

import "fmt"

const boilingF = 212.0 //定义一个常量，常量一旦定义，其值在程序运行过程中不能被修改

func main() {
	var f = boilingF                      //常量boilingF赋给变量f
	var c = (f - 32) * 5 / 9              //将f转换为摄氏度
	fmt.Printf("沸点是 = %g°F或%g°C\n", f, c) //%g是占位符
}
