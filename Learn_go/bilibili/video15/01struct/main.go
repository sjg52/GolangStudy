package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"名字",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int)
}
