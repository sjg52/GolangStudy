package main

import "fmt"

func main() {
	//	var ch1 chan int        //通道是一种引用类型，需要初始化之后才能使用
	//	ch1 = make(chan int, 1) //1代表缓存区的大小，1是指缓冲区放一个值，这种是带缓冲区通道，如果不加1那就是无缓冲区通道

	ch1 := make(chan int, 1) //上面的是分了两步声明，也可以像这样一行声明
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
	close(ch1) //关闭通道
}
