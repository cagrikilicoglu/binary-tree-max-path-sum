package tree

import (
	"binary-tree-max-path-sum/internal/api"
	"log"
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

func ResponseToBinaryTree(t api.Tree) *Node {
	tree := newBinaryTree()
	list := []*Node{}

	for i := range t.Nodes {
		node := responseToNode(*t.Nodes[i])
		list = append(list, node)
	}

	for i := 0; i < len(t.Nodes); i++ {

		if *t.Root == *t.Nodes[i].ID {
			tree.Root = list[i]
			log.Println("root in if", tree.Root)
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

// func responseToNode(nr NodeResponse) *Node {
// 	node := newNode()
// 	node.Value = nr.Value
// 	return node
// }
func MaxSum(n *Node) int {
	maxSum := math.MinInt
	PathSum(n, &maxSum)
	return maxSum
}
func PathSum(n *Node, maxSum *int) int {

	if n == nil {
		return math.MinInt
	}

	left := PathSum(n.Left, maxSum)
	right := PathSum(n.Right, maxSum)
	var currentMax int

	// to prevent integer overlow
	if n.Value < 0 && (right == math.MinInt || left == math.MinInt) {

		currentMax = n.Value
		*maxSum = max(*maxSum, currentMax)
	} else {

		currentMax = max(n.Value, max(n.Value+right, n.Value+left))
		*maxSum = max(*maxSum, max(n.Value+right+left, currentMax))
	}

	return currentMax

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// var tree = BinaryTree{}

// func responseToNode

// func responseToBinaryTree(br BinaryTreeResponse) *Node {
// currentNode = responseToNode(br.Nodes[0])

// if len(br.Nodes) == 0 {
// 	return nil
// }

// br.Nodes = br.Nodes[1:]
// currentNode.right = responseToBinaryTree

// for i := range br.Nodes {
// 	node := newNode()
// 	node.value = br.Nodes[i].Value
// 	if br.Nodes[i].Id == br.root {

// 	}
// }
// return constructNodes(br.Nodes)

// }
