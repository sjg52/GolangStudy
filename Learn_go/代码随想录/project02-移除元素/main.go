package main

import "fmt"

func removeElement(nums []int, val int) int {
	slow := 0 //定义慢指针
	//定义快指针遍历nums
	for _, fast := range nums {
		if fast != val {
			//如果不等于就往前挪
			nums[slow] = fast
			//然后去下一位
			slow++
		}
	}
	//最后返回慢指针
	return slow
}
func main() {
	//你的函数函数应该返回 k = 2, 并且 nums 中的前两个元素均为 2。
	//你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
