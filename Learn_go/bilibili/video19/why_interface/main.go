package main

import "fmt"

// 为什么需要接口
type dog struct{}

func (d dog) say() {
	fmt.Println("汪汪汪！")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("喵喵喵！")
}

// 接口不管你是什么类型，只管要实现什么方法
// 定义一个类型，一个抽象的类型，只要实现了say()这个方法的类型都可以称为sayer类型
type sayer interface {
	say()
}

// 动作函数
func rua(arg sayer) {
	arg.say() // 不管传进来什么都执行say
}

func main() {
	c1 := cat{}
	rua(c1)
	d1 := dog{}
	rua(d1)
}
