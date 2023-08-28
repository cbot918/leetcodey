package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var root *TreeNode

func Insert(val int) {
	newNode := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	if root == nil {
		root = newNode
		return
	}
	insert(root, newNode)
}
func insert(root, node *TreeNode) {
	if node.Val < root.Val {
		if root.Left != nil {
			insert(root.Left, node)
			return
		}
		root.Left = node
	} else {
		if root.Right != nil {
			insert(root.Right, node)
			return
		}
		root.Right = node
	}
}

func DFS(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	DFS(root.Left)
	DFS(root.Right)
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return SearchBST(root.Left, val)
	} else {
		return SearchBST(root.Right, val)
	}
}

func main() {
	nums := []int{18, 2, 22, 63, 84}

	for _, num := range nums {
		Insert(num)
	}

	// DFS(root)
	result := SearchBST(root, 22)
	fmt.Println("result: ", result)
}
