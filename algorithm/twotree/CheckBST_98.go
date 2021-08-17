package main

import (
	"fmt"
	"math"
)

/**
BST与其验证， lc 98
*/
func main() {
	tree := buildBSTTree()
	bst := checkBST(tree)
	fmt.Println("isBST =", bst)
}

func checkBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.data || root.data >= max {
		return false
	}
	return isBST(root.left, min, root.data) && isBST(root.right, root.data, max)
}

func buildBSTTree() *TreeNode {
	nodeF := &TreeNode{nil, nil, 7}
	nodeG := &TreeNode{nil, nil, 11}
	nodeB := &TreeNode{nil, nil, 3}
	nodeC := &TreeNode{nodeF, nodeG, 8}
	nodeA := &TreeNode{nodeB, nodeC, 5}
	return nodeA
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}
