package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	ans := sortedSquares(nums)
	fmt.Println(ans)
}

func sortedSquares(nums []int) []int {
	//切片的长度
	n := len(nums)
	//初始化指针，i指向切片起始，j指向切片末尾，k指向结果切片的末尾位置
	i, j, k := 0, n-1, n-1
	//创建个长度为n的切片用于存储最终结果
	ans := make([]int, n)
	//当i小于等于j的时候，继续循环
	for i <= j {
		//计算i位置元素的平方
		lm := nums[i] * nums[i]
		//计算j位置元素的平方
		rm := nums[j] * nums[j]
		//比较lm和rm的大小
		if lm > rm {
			//如果lm大于rm，则lm放入k位置
			ans[k] = lm
			//然后i右移
			i++
		} else {
			//否则rm放入k位置
			ans[k] = rm
			//然后j左移
			j--
		}
		//循环一次之后k左移
		k--
	}
	//返回结果切片
	return ans
}
