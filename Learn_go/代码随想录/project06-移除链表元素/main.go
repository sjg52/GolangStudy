package main

import "fmt"

func main() {
	// 创建链表 [1,2,6,3,4,5,6]
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 6}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 6}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	val := 6
	newHead := removeElements(head, val)

	// 打印原始链表
	fmt.Print("原始链表: ")
	printList(head)

	// 打印删除后的链表
	fmt.Print("删除后的链表: ")
	printList(newHead)
}

// 辅助函数：打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 如果头指针就是满足条件的节点就直接指向下一个
	for head != nil && head.Val == val {
		head = head.Next
	}

	//定义一个指针指向head
	cur := head

	for cur != nil && cur.Next != nil {
		//如果下一个指针指向的是val就指向下下个
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
