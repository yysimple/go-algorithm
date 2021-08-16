package main

import "fmt"

/**
层次遍历与BFS， lc 102
*/
func main() {
	tree := buildTwoTree()
	order := levelOrder(tree)
	for key, value := range order {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}
}

func levelOrder(root *TreeNode) map[int][]*TreeNode {
	res := make(map[int][]*TreeNode)
	level := 1
	if root == nil {
		return map[int][]*TreeNode{level: {&TreeNode{}}}
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		queueLen := len(queue)
		res = map[int][]*TreeNode{level: queue}
		for queueLen > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			queueLen--
		}
		level++
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func buildTwoTree() *TreeNode {
	nodeH := &TreeNode{nil, nil, "H"}
	nodeD := &TreeNode{nil, nil, "D"}
	nodeE := &TreeNode{nil, nil, "E"}
	nodeF := &TreeNode{nil, nil, "F"}
	nodeG := &TreeNode{nil, nodeH, "G"}
	nodeB := &TreeNode{nodeD, nodeE, "B"}
	nodeC := &TreeNode{nodeF, nodeG, "C"}
	nodeA := &TreeNode{nodeB, nodeC, "A"}
	return nodeA
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  string
}
