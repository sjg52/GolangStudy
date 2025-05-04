package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	// 获取数组的长度
	n := len(nums)
	// 滑动窗口左边界的指针
	left := 0
	// 当前滑动窗口内元素相加的和
	sum := 0
	// 记录最小的长度
	minLength := math.MaxInt32

	// 循环遍历确认右指针
	for right := 0; right < n; right++ {
		sum += nums[right]
		//如果和大于target就缩小查找
		for sum >= target {
			// 计算长度
			length := right - left + 1
			// 如果发现更小的就更新一下
			if length < minLength {
				minLength = length
			}
			sum -= nums[left]
			// 左指针右移
			left++
		}

	}
	// 如果找不到就说明没有，返回0
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}
