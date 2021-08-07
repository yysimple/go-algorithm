package main

import "fmt"

/**
两数相加，lc 2
*/
func main() {
	l1 := &ListNode{
		&ListNode{
			&ListNode{
				nil,
				3},
			4},
		2}
	l2 := &ListNode{
		&ListNode{
			&ListNode{
				nil,
				4},
			6},
		5}
	add := twoNumAdd(l1, l2)
	//fmt.Println("add =", add)
	printNode(add)
}

/**
这里很简单，就是两个链表同时往后遍历，然后将值赋给tmp中间变量，然后两个链表的值相加完之后，赋值给需要返回的链表中
*/
func twoNumAdd(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head
	tmp := 0
	for l1 != nil && l2 != nil {
		if l1 != nil {
			tmp += l1.data
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.data
			l2 = l2.Next
		}
		head.Next = &ListNode{nil, tmp % 10}
		tmp = tmp / 10
		head = head.Next
	}
	return res.Next
}

type ListNode struct {
	Next *ListNode
	data int
}

func printNode(node *ListNode) {
	var nodeData []int
	res := &ListNode{}
	res = node
	for res.Next != nil {
		nodeData = append(nodeData, res.data)
		res = res.Next
	}
	nodeData = append(nodeData, res.data)
	fmt.Println("nodeData =", nodeData)
}
