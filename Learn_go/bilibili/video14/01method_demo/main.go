package main

import "fmt"

// 方法的定义实例

// Person试一个结构体
type Person struct {
	name string
	age  int8
}

// NewPerson是一个Person类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream是为Person类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是！\n", p.name)
}

// SetAge是修改年龄的方法
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2是一个使用值接收者来修改年龄的方法
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("路人甲", int8(18))
	//(*p1).Dream()
	//p1.Dream()

	// 调用修改年龄的方法
	fmt.Println(p1.age)
	p1.SetAge(int8(20))
	fmt.Println(p1.age)

	p1.SetAge2(int8(30))
	fmt.Println(p1.age) //20?
}
