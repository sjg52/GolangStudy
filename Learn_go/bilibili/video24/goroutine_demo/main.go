package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello 开拓者", i)
	wg.Done() //通知wg把计数器-1
}
func main() { //开启一个主goroutine去执行main函数
	wg.Add(10)                //计数牌+8
	for i := 0; i < 10; i++ { //计数牌+1
		go hello(i) //开启了一个goroutine去执行hello这个函数
	}

	fmt.Println("hello 审神者")
	//time.Sleep(time.Second)
	wg.Wait() //阻塞，等所有小弟都干完活之后才收兵
}
