package main

import "fmt"

func main() {
	n := 4
	res := generateMatrix(n)
	fmt.Println(res)
}

func generateMatrix(n int) [][]int {
	// 初始化横坐标，纵坐标，和终止标志
	startx, starty := 0, 0
	// 需要转n/2圈
	loop := n / 2
	// 当n为奇数的时候中心的坐标
	center := n / 2
	// 偏移量，控制每一圈的边界
	offset := 1
	// 起始填充的数字，从1开始
	count := 1
	// 保存最后的矩阵
	res := make([][]int, n)
	// 初始化每行
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// 循环填充每圈
	for loop > 0 {
		// 初始化当前这一圈的起始坐标
		i, j := startx, starty

		//从左到右填充上边界
		// 行数不变（startx），列数从 starty 增加到 n - offset - 1
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}

		//从上到下填充右边界
		// 列数不变（j），行数从 startx 增加到 n - offset - 1
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}

		//从右到左填充下边界
		// 行数不变（i），列数从 j 减少到 starty + 1
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}

		//从下到上填充左边界
		// 列数不变（j），行数从 i 减少到 startx + 1
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}

		//新的一圈更新起始坐标
		startx++
		starty++
		offset++
		loop--
	}
	// 如果n是奇数，中心单独填充
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
