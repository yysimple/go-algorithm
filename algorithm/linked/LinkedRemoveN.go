package main

import "fmt"

/**
删除链表倒数第N个节点，lc 19
*/
func main() {
	head := &ListNode{&ListNode{&ListNode{&ListNode{&ListNode{nil, 0}, 4}, 3}, 2}, 1}
	// 倒数第一个元素，这里需要打印数据，所以需要多加一个节点，最后节点为空节点，所以需要
	var n int = 1
	n = n + 1
	end := removeNthFromEnd(head, n)
	for end.Next != nil {
		fmt.Println("node =", end)
		end = end.Next
	}
}

type ListNode struct {
	Next *ListNode
	data int
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 定义一个空节点，相当于
	result := &ListNode{}
	// 将空节点作为头节点
	result.Next = head
	// 定义一个pre节点
	var pre *ListNode
	i := 1
	delete := result
	/**
		i = 1, head.data = 1, pre={}, delete = {}
		i = 2, head.data = 2, pre={}, delete = {}
	    i = 3, head.data = 3, pre={}, delete = 1
		i = 4, head.data = 4, pre=1, delete = 2
		i = 5, head.data = nil, pre=2, delete = 3
		i = 6, break
	*/
	for head != nil {
		if i >= n {
			pre = delete
			delete = delete.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}
