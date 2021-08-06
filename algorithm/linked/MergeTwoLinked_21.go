package main

import "fmt"

/**
合并两个链表，lc 21
*/
func main() {
	l1 := &ListNode{
		&ListNode{
			&ListNode{
				nil,
				3},
			2},
		1}
	l2 := &ListNode{
		&ListNode{
			&ListNode{
				nil,
				4},
			2},
		1}

	linked := mergeTwoLinked(l1, l2)
	printNode(linked)
}

func mergeTwoLinked(l1 *ListNode, l2 *ListNode) *ListNode {
	virtualNode := &ListNode{}
	result := virtualNode
	for l1 != nil && l2 != nil {
		if l1.data < l2.data {
			virtualNode.Next = l1
			l1 = l1.Next
		} else {
			virtualNode.Next = l2
			l2 = l2.Next
		}
		virtualNode = virtualNode.Next
	}
	if l1 != nil {
		virtualNode.Next = l1
	}
	if l2 != nil {
		virtualNode.Next = l2
	}
	return result.Next
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
