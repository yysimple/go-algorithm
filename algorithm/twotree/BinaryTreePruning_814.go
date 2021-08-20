package main

import "fmt"

/**
二叉树的剪枝 lc， 814
*/
func main() {
	tree := buildTwoTree()
	res := printTree(tree)
	fmt.Println("print tree =", res)
	node := pruneTree(tree)
	nodeRes := printTree(node)
	fmt.Println("nodeRes =", nodeRes)

}

func pruneTree(root *TreeNode) *TreeNode {
	return deal(root)
}

func deal(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.left = deal(node.left)
	node.right = deal(node.right)
	// 只要没有左右节点，然后自己本身等于0的，则删除
	if node.left == nil && node.right == nil && node.data == 0 {
		return nil
	}
	return node
}

func printTree(root *TreeNode) [][]int {
	var res [][]int
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
				res = append(res, []int{node.data})
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

func buildTwoTree() *TreeNode {
	nodeI := &TreeNode{nil, nil, 0}
	nodeF := &TreeNode{nil, nil, 1}
	nodeG := &TreeNode{nil, nil, 0}
	nodeD := &TreeNode{nil, nil, 1}
	nodeH := &TreeNode{nodeI, nil, 1}
	nodeC := &TreeNode{nodeG, nodeF, 0}
	nodeB := &TreeNode{nodeH, nodeD, 1}
	nodeA := &TreeNode{nodeB, nodeC, 1}
	return nodeA
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}
