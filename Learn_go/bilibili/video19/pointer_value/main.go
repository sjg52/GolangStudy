package main

import "fmt"

//使用值接收者实现接口和使用指针接收实现接口的区别

type mover interface {
	move()
}

type person struct {
	name string
	age  int8
}

//使用值接受者实现接口
func (p person) move() {
	fmt.Printf("%s在跑……\n", p.name)
}

func main() {
	var m mover
	p1 := person{ //p1是person类型的值
		name: "路人A",
		age:  18,
	}
	p2 := &person{ //p2是person类型的指针
		name: "路人B",
		age:  16,
	}
	m = p1
	m.move()
	m = p2
	m.move()
	fmt.Println(m)
}
