package main

import "fmt"

/**
BST 的查找，lc 700
*/
func main() {
	tree := buildBSTSearchTree()
	subTree := searchSubTree(tree, 8)
	fmt.Println("subTree =", subTree.data)
	fmt.Println("subTree.left =", subTree.left.data)
	fmt.Println("subTree.right =", subTree.right.data)
	fmt.Println("---------------------------")
	search := dgSearch(tree, 8)
	fmt.Println("search =", search.data)
	fmt.Println("search.left =", search.left.data)
	fmt.Println("search.right =", search.right.data)
}

/**
迭代的方式
*/
func searchSubTree(root *TreeNode, target int) *TreeNode {
	for root != nil {
		if root.data == target {
			return root
		} else if root.data > target {
			root = root.left
		} else {
			root = root.right
		}

	}
	return nil
}

/**
迭代的方式
*/
func dgSearch(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.data > target {
		return dgSearch(root.left, target)
	} else if root.data < target {
		return dgSearch(root.right, target)
	} else {
		return root
	}
}

func buildBSTSearchTree() *TreeNode {
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
