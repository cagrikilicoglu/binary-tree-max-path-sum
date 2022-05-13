package tree

import "binary-tree-max-path-sum/internal/api"

func responseToNode(an api.Node) *Node {
	node := newNode()
	node.Value = int(*an.Value)
	return node
}
