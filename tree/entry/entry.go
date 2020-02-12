package main

import (
	"eg/egolang/tree"
	"fmt"
)

func main() {
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)
	root.Print()
	fmt.Println()

	root.Traverse()
}
