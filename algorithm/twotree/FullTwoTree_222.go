package main

import "fmt"

/**
完全二叉树,lc 222
*/
func main() {
	fmt.Println("sss =", 1<<2)
	tree := buildBalanceTree()
	nodes := countNodesDg(tree)
	fmt.Println("node num =", nodes)
	i := countNodes(tree)
	fmt.Println("no dg = ", i)
	dg := countNodesAndDg(tree)
	fmt.Println("dg =", dg)
}

func countNodesDg(root *TreeNode) int {
	if root != nil {
		return 1 + countNodesDg(root.right) + countNodesDg(root.left)
	}
	return 0
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := totalLevel(root.left)
	right := totalLevel(root.right)
	if left == right {
		return 1 + (1<<left - 1) + (1<<right - 1)
	} else {
		return countNodesDg(root.left) + (1 << right)
	}
}

func countNodesAndDg(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := totalLevel(root.left)
	right := totalLevel(root.right)
	if left == right {
		return countNodesAndDg(root.right) + (1 << left)
	} else {
		return countNodesAndDg(root.left) + (1 << right)
	}
}

/**
这里只考虑左节点的情况，完全二叉树，只会往左填充的，所以不管非叶子节点的子节点，一定是左节点是最后一位
*/
func totalLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.left
	}
	return level
}

func buildBalanceTree() *TreeNode {
	nodeF := &TreeNode{nil, nil, 7}
	node10 := &TreeNode{nil, nil, 15}
	node12 := &TreeNode{nil, nil, 14}
	//nodeG := &TreeNode{node10, node12, 11}
	nodeG := &TreeNode{nil, nil, 11}
	nodeB := &TreeNode{node10, node12, 3}
	nodeC := &TreeNode{nodeF, nodeG, 8}
	nodeA := &TreeNode{nodeB, nodeC, 5}
	return nodeA
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}
