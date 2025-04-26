package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() { //开启一个主goroutine去执行main函数
	wg.Add(10)                //计数牌+8
	for i := 0; i < 10; i++ { //计数牌+1
		go func(i int) {
			fmt.Println("hello 开拓者", i)
			wg.Done()
		}(i)
	}

	fmt.Println("hello 审神者")
	//time.Sleep(time.Second)
	wg.Wait() //阻塞，等所有小弟都干完活之后才收兵
}
