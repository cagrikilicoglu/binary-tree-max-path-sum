package tree

import (
	"binary-tree-max-path-sum/internal/api"
	"math"
)

type BinaryTree struct {
	Root *Node
}

func newBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func newNode() *Node {
	return &Node{}
}

// CreateBinaryTree creates a BinaryTree entity from a given input binary tree in the form of json
func CreateBinaryTree(t api.Tree) *Node {
	tree := newBinaryTree()
	list := []*Node{}
	if len(t.Nodes) == 0 {
		return nil
	}
	for i := range t.Nodes {
		node := inputToNode(*t.Nodes[i])
		list = append(list, node)
	}

	for i := 0; i < len(t.Nodes); i++ {

		if t.Root == *t.Nodes[i].ID {
			tree.Root = list[i]
		}

		for j := i + 1; j < len(t.Nodes); j++ {
			if t.Nodes[i].Right == *t.Nodes[j].ID {
				list[i].Right = list[j]
			}
			if t.Nodes[i].Left == *t.Nodes[j].ID {
				list[i].Left = list[j]
			}
		}
	}
	return tree.Root
}

// MaxSum calculates the maximum sum of all paths of a binary tree
func MaxSum(n *Node) int {
	if n == nil {
		return 0
	}
	maxSum := math.MinInt
	PathSum(n, &maxSum)
	return maxSum
}

// PathSum calculates the maximum sum of paths passing through each node recursively
func PathSum(n *Node, maxSum *int) int {

	if n == nil {
		return math.MinInt
	}

	left := PathSum(n.Left, maxSum)
	right := PathSum(n.Right, maxSum)
	var currentMax int

	// to prevent integer overlow when the value is negative
	if n.Value < 0 && (right == math.MinInt || left == math.MinInt) {
		currentMax = n.Value
		*maxSum = max(*maxSum, currentMax)
	} else {
		currentMax = max(n.Value, max(n.Value+right, n.Value+left))
		*maxSum = max(*maxSum, max(n.Value+right+left, currentMax))
	}

	return currentMax

}

// max calculates the maximum value of two given inputs
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// inputToNode creates a node entity from given input noed in the form of json and sets its value to input node value
func inputToNode(an api.Node) *Node {
	node := newNode()
	node.Value = int(*an.Value)
	return node
}
