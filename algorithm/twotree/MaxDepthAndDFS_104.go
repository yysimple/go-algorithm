package main

import "fmt"

/**
最大深度与DFS， lc 104
*/
func main() {
	tree := buildTwoTree()
	depth := maxDepthHeight(tree)
	fmt.Println("depthHeight =", depth)
	width := maxDepthWidth(tree)
	fmt.Println("depthWidth =", width)
}

/**
广度遍历
*/
func maxDepthWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			// 截取取下一个数据，相当于遍历队列，
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			// 当前行的节点遍历完成后，最后会得到下一层节点，之后没有下一层，即 len(queue) = 0
			sz--
		}
		depth++
	}
	return depth
}

/**
递归查询最大的深度：深度优先搜索
*/
func maxDepthHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthHeight(root.left), maxDepthHeight(root.right)) + 1
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
