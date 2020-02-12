package tree

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node *TreeNode) setValue(value int) {
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}
