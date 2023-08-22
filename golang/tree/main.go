package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Insert(num int) {
	newNode := &Node{Val: num}
	if root == nil {
		root = newNode
		return
	}
	insert(root, newNode)
}
func insert(root, newNode *Node) {
	if newNode.Val < root.Val {
		if root.Left != nil {
			insert(root.Left, newNode)
			return
		}
		root.Left = newNode
	} else {
		if root.Right != nil {
			insert(root.Right, newNode)
			return
		}
		root.Right = newNode
	}
}
func PreOrder(root *Node) {
	if root == nil {
		// logf("nil ")
	} else {
		logf("%d ", root.Val)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}
func InOrder(root *Node) {
	if root == nil {
		// logf("nil ")
	} else {
		InOrder(root.Left)
		logf("%d ", root.Val)
		InOrder(root.Right)
	}
}
func PostOrder(root *Node) {
	if root == nil {
		// logf("nil ")
	} else {
		PostOrder(root.Left)
		PostOrder(root.Right)
		logf("%d ", root.Val)
	}
}
func TreeHeight(root *Node) int {
	if root == nil {
		return 0
	} else {
		left := TreeHeight(root.Left)
		right := TreeHeight(root.Right)

		if left > right {
			return left + 1
		} else {
			return right + 1
		}

	}
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*Node{}
	q = append(q, root)

	ans := [][]int{}
	for len(q) > 0 {

		temp := []int{}
		i := 0
		lenq := len(q)
		for i = 0; i < lenq; i++ {
			curNode := q[i]
			temp = append(temp, curNode.Val)
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}

		log(temp)
		ans = append(ans, temp)
		q = q[i:]
	}
	return ans
}

var root *Node
var log = fmt.Println
var logf = fmt.Printf

func main() {
	// nums := []int{5, 2, 4, 7, 6}
	nums := []int{5, 2, 7, 6, 8}

	for _, num := range nums {
		Insert(num)
	}

	// PreOrder(root)
	// InOrder(root)
	PostOrder(root)

	log("\nTreeHeight:")
	log(TreeHeight(root))

	log("\nlevelOrder:")
	log(levelOrder(root))

}
