package tree

import (
	"log"
	"math"
)

type BinaryTreeResponse struct {
	Nodes []NodeResponse `json:"nodes"`
	Root  string
}

type NodeResponse struct {
	Id    string `json:"id"`
	Left  string `json:"left"`
	Right string `json:"right"`
	Value int    `json:"value"`
}

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

func ResponseToBinaryTree(br BinaryTreeResponse) *BinaryTree {
	tree := newBinaryTree()
	list := []*Node{}

	for i := range br.Nodes {
		node := responseToNode(br.Nodes[i])
		list = append(list, node)
	}

	for i := 0; i < len(br.Nodes); i++ {
		if br.Root == br.Nodes[i].Id {
			tree.Root = list[i]
		}
		for j := i + 1; j < len(br.Nodes); j++ {
			if br.Nodes[i].Right == br.Nodes[j].Id {
				list[i].Right = list[j]
			}
			if br.Nodes[i].Left == br.Nodes[j].Id {
				list[i].Left = list[j]
			}

		}
	}
	return tree

}

func responseToNode(nr NodeResponse) *Node {
	node := newNode()
	node.Value = nr.Value
	return node
}
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
	log.Println("val", n.Value)
	log.Println("max", *maxSum)
	return currentMax

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
