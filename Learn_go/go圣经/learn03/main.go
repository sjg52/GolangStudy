package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "省略末尾的换行符") //如果使用了 -n 标志，则不会在输出末尾添加换行符
var sep = flag.String("s", " ", "分隔符")    //如果使用了 -s 标志，可以指定不同的分隔符

// 它接受命令行参数，并使用指定的分隔符（默认为空格）将它们连接起来
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
