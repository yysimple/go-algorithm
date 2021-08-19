package main

import "fmt"

/**
平衡二叉树,lc 110
*/
func main() {
	tree := buildBalanceTree()
	balance := isBalance(tree)
	fmt.Println("is balance? ", balance)
}

func isBalance(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 也需要判断子树也是平衡树
	if !isBalance(root.left) || !isBalance(root.right) {
		return false
	}
	depthLeft := maxDepth(root.left) + 1
	depthRight := maxDepth(root.right) + 1
	if abs(depthRight-depthLeft) > 1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.left), maxDepth(root.right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func buildBalanceTree() *TreeNode {
	nodeF := &TreeNode{nil, nil, 7}
	//node10 := &TreeNode{nil, nil, 11}
	//node12 := &TreeNode{nil, nil, 11}
	//nodeG := &TreeNode{node10, node12, 11}
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
