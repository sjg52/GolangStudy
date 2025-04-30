package main

import "fmt"

// search performs a binary search for target in a sorted slice of integers and returns its index, or -1 if not found.
func search(nums []int, target int) int {
	left := 0              //初始化左边界
	right := len(nums) - 1 //初始化右边界
	//查找范围
	for left <= right {
		//取中间值，可能会存在溢出问题
		middle := (left + right) / 2
		if target > nums[middle] {
			//如果大于中间值
			//往右查询
			left = middle + 1
		} else if target < nums[middle] {
			//如果小于中间值
			//往左查询
			right = middle - 1
		} else {
			//否则等于中间值
			return middle
		}
	}
	//都没有返回-1表示不在数组内
	return -1
}

//左闭右开区间
/*
func search(nums []int, target int) int {
	left := 0          //初始化左边界
	right := len(nums) //初始化右边界
	//查找范围
	for left < right {
		//取中间值，不会存在溢出问题，因为先计算 right - left，这个差值通常不会超出数据类型的表示范围，然后再进行右移和加法操作，也不会产生溢出
		//>>是右移运算符，它属于位运算符的一种。右移运算符会把一个数的二进制位向右移动指定的位数
		mid := left + (right-left)>>1
		// 根据 nums[mid] 和 target 的大小关系
		// 调整区间范围
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}

	}
}
*/

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	//nums := []int{-1,0,3,5,9,12}
	var target = 9
	//var target = 2
	fmt.Println(search(nums, target))
}
