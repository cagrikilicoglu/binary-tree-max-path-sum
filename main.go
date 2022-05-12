package main

import (
	"binary-tree-max-path-sum/internal/tree"
	"fmt"
)

func main() {
	// myTree := tree.BinaryTreeResponse{
	// 	Nodes: []tree.NodeResponse{
	// 		{Id: "1", Left: "2", Right: "3", Value: 1},
	// 		{Id: "3", Left: "6", Right: "7", Value: 3},
	// 		{Id: "7", Left: "null", Right: "null", Value: 7},
	// 		{Id: "6", Left: "null", Right: "null", Value: 6},
	// 		{Id: "2", Left: "4", Right: "5", Value: 2},
	// 		{Id: "5", Left: "null", Right: "null", Value: 5},
	// 		{Id: "4", Left: "null", Right: "null", Value: 4},
	// 	},
	// 	Root: "1",
	// }
	myTree2 := tree.BinaryTreeResponse{
		Nodes: []tree.NodeResponse{
			{Id: "1", Left: "-10", Right: "-5", Value: 1},
			{Id: "-5", Left: "-20", Right: "-21", Value: -5},
			{Id: "-21", Left: "100-2", Right: "1-3", Value: -21},
			{Id: "1-3", Left: "null", Right: "null", Value: 1},
			{Id: "100-2", Left: "null", Right: "null", Value: 100},
			{Id: "-20", Left: "100", Right: "2", Value: -20},
			{Id: "2", Left: "null", Right: "null", Value: 2},
			{Id: "100", Left: "null", Right: "null", Value: 100},
			{Id: "-10", Left: "30", Right: "45", Value: -10},
			{Id: "45", Left: "3", Right: "-3", Value: 45},
			{Id: "-3", Left: "null", Right: "null", Value: -3},
			{Id: "3", Left: "null", Right: "null", Value: 3},
			{Id: "30", Left: "5", Right: "1-2", Value: 30},
			{Id: "1-2", Left: "null", Right: "null", Value: 1},
			{Id: "5", Left: "null", Right: "null", Value: 5},
		},
		Root: "1",
	}

	// myBinaryTree := tree.ResponseToBinaryTree(myTree)
	// fmt.Println(*myBinaryTree.Root)
	// fmt.Println(*myBinaryTree.Root.Right)
	// fmt.Println(*myBinaryTree.Root.Left)
	// fmt.Println(*myBinaryTree.Root.Right.Right)
	// fmt.Println(*myBinaryTree.Root.Right.Left)
	// fmt.Println(*myBinaryTree.Root.Left.Right)
	// fmt.Println(*myBinaryTree.Root.Left.Left)
	// fmt.Println(tree.MaxSum(myBinaryTree.Root))

	myBinaryTree2 := tree.ResponseToBinaryTree(myTree2)
	fmt.Println(*myBinaryTree2.Root)
	fmt.Println(*myBinaryTree2.Root.Right)
	fmt.Println(*myBinaryTree2.Root.Left)
	fmt.Println(*myBinaryTree2.Root.Right.Right)
	fmt.Println(*myBinaryTree2.Root.Right.Left)
	fmt.Println(*myBinaryTree2.Root.Left.Right)
	fmt.Println(*myBinaryTree2.Root.Left.Left)
	fmt.Println(*myBinaryTree2.Root.Right.Right.Right)
	fmt.Println(*myBinaryTree2.Root.Right.Right.Left)
	fmt.Println(*myBinaryTree2.Root.Right.Left.Right)
	fmt.Println(*myBinaryTree2.Root.Right.Left.Left)
	fmt.Println(*myBinaryTree2.Root.Left.Right.Right)
	fmt.Println(*myBinaryTree2.Root.Left.Right.Left)
	fmt.Println(*myBinaryTree2.Root.Left.Left.Right)
	fmt.Println(*myBinaryTree2.Root.Left.Left.Left)
	fmt.Println(tree.MaxSum(myBinaryTree2.Root))
}
