package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func NewTree() *TreeNode {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	return root
}

// preOrderTraverse 前序遍历
func preOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Printf("%d ", tree.Val)
	preOrderTraverse(tree.Left)
	preOrderTraverse(tree.Right)
}

// preOrderTraverseNoRec 前序遍历 非递归
func preOrderTraverseNoRec(tree *TreeNode) {
	stack := list.New()
	root := tree
	for root != nil || stack.Len() != 0 {
		for root != nil {
			fmt.Printf("%d ", root.Val)
			stack.PushBack(root)
			root = root.Left
		}

		if stack.Len() != 0 {
			v := stack.Back()
			root = v.Value.(*TreeNode)
			root = root.Right
			stack.Remove(v)
		}
	}
}

func midOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	midOrderTraverse(tree.Left)
	fmt.Printf("%d ", tree.Val)
	midOrderTraverse(tree.Right)
}

func midOrderTraverseNoRec(tree *TreeNode) {
	if tree == nil {
		return
	}
	root := tree
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}

		if stack.Len() != 0 {
			elem := stack.Back()
			node := elem.Value.(*TreeNode)
			fmt.Printf("%d ", node.Val)
			stack.Remove(elem)
			root = node.Right
		}
	}

}

func postOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	postOrderTraverse(tree.Left)
	postOrderTraverse(tree.Right)
	fmt.Printf("%d ", tree.Val)
}

func postOrderTraverseNoRec(tree *TreeNode) {
	if tree == nil {
		return
	}
	t := tree
	stack := list.New()
	var preVisited *TreeNode

	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}

		v := stack.Back()
		top := v.Value.(*TreeNode)

		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right {
			fmt.Printf("%d ", top.Val)
			preVisited = top
			stack.Remove(v)
		} else {
			t = top.Right
		}
	}
}

func main() {
	tree := NewTree()
	preOrderTraverse(tree)
	fmt.Println()
	preOrderTraverseNoRec(tree)
	fmt.Println("\n==================")
	midOrderTraverse(tree)
	fmt.Println()
	midOrderTraverseNoRec(tree)
	fmt.Println("\n==================")

	postOrderTraverse(tree)
	fmt.Println()
	postOrderTraverseNoRec(tree)
}
