package main

import "fmt"

/**
层次遍历与BFS， lc 102
*/
func main() {
	tree := buildTwoTree()
	order := levelOrder(tree)
	fmt.Println("order =", order)

	dg := levelOrderDg(tree)
	fmt.Println("dg =", dg)
}

/**
广度优先搜索
*/
func levelOrder(root *TreeNode) [][]string {
	var res [][]string
	level := 0
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		queueLen := len(queue)
		for queueLen > 0 {
			node := queue[0]
			if len(res) == level {
				res = append(res, []string{node.data})
			} else {
				res[level] = append(res[level], node.data)
			}
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

func levelOrderDg(root *TreeNode) [][]string {
	return dfs(root, 0, [][]string{})
}

func dfs(root *TreeNode, level int, res [][]string) [][]string {
	if root == nil {
		return res
	}
	if len(res) == level {
		res = append(res, []string{root.data})
	} else {
		res[level] = append(res[level], root.data)
	}
	res = dfs(root.left, level+1, res)
	res = dfs(root.right, level+1, res)
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
