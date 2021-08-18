package main

import "fmt"

/**
BST 的删除, lc 450
*/
func main() {

	tree := buildBSTRemoveTree()
	node := deleteNodeAfter(tree, 5)
	fmt.Println("node =", node.data)
	fmt.Println("node =", node.right.data)
	fmt.Println("node =", node.left.data)
	fmt.Println("---------------------------")
	before := deleteNodeBefore(tree, 5)
	fmt.Println("before =", before.data)
	fmt.Println("before =", before.right.data)
	fmt.Println("before =", before.left.data)
}

func deleteNodeAfter(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.data {
		root.left = deleteNodeAfter(root.left, key)
		return root
	}
	if key > root.data {
		root.right = deleteNodeAfter(root.right, key)
		return root
	}
	//到这里意味已经查找到目标
	if root.right == nil {
		//右子树为空
		return root.left
	}
	if root.left == nil {
		//左子树为空
		return root.right
	}
	// 将后一个节点赋值
	minNode := root.right
	for minNode.left != nil {
		//查找后继
		minNode = minNode.left
	}
	root.data = minNode.data
	// 将其后一个节点
	root.right = deleteMinNode(root.right)
	return root
}

func deleteMinNode(root *TreeNode) *TreeNode {
	if root.left == nil {
		pright := root.right
		root.right = nil
		return pright
	}
	root.left = deleteMinNode(root.left)
	return root
}

func deleteNodeBefore(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.data > target {
		root.left = deleteNodeBefore(root.left, target)
		return root
	}
	if root.data < target {
		root.right = deleteNodeBefore(root.right, target)
		return root
	}
	if root.left == nil {
		return root.right
	}
	if root.right == nil {
		return root.left
	}
	maxNode := root.left
	for maxNode.right != nil {
		maxNode = maxNode.right
	}
	root.data = maxNode.data
	root.left = deleteMaxNode(root.left)
	return root
}

func deleteMaxNode(node *TreeNode) *TreeNode {
	if node.right != nil {
		maxLeft := node.left
		node.left = nil
		return maxLeft
	}
	node.right = deleteMaxNode(node.right)
	return node
}

func buildBSTRemoveTree() *TreeNode {
	node4 := &TreeNode{nil, nil, 4}
	node2 := &TreeNode{nil, nil, 2}
	node9 := &TreeNode{nil, nil, 9}
	node6 := &TreeNode{nil, nil, 6}
	node7 := &TreeNode{node6, node9, 7}
	node3 := &TreeNode{node2, node4, 3}
	node5 := &TreeNode{node3, node7, 5}
	return node5
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}
