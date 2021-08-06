package main

import "fmt"

/**
环形链表，lc 141
*/
func main() {
	l5 := &ListNode{nil, 5}
	l4 := &ListNode{l5, 4}
	l3 := &ListNode{l4, 3}
	l2 := &ListNode{l3, 2}
	l1 := &ListNode{l2, 1}
	l5.Next = l2

	cycle := hasCycle(l1)
	fmt.Println("cycle =", cycle)

	point := fastSlowPoint(l1)
	fmt.Println("point =", point)
}

/**
这里就是很简单，将每个node都放入到 map的key中，值不重要，然后遍历，出现过两次后，则证明有环
*/
func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]string{}
	for head != nil {
		k, s := seen[head]
		fmt.Println("k =", k, "s =", s)
		//
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = "hahahahha"
		head = head.Next
	}
	return false
}

/**
快慢指针法,快的先走一步，只要快慢还能重合，那证明就又环形
*/
func fastSlowPoint(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		head = head.Next
		fast = fast.Next.Next
	}
	return false
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
