package main

import (
	"fmt"
	"github.com/xinwangqilin/base_learn/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding 内嵌 语法糖
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	// 建立了一个根节点
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Node.Traverse()
	root.Traverse()
	fmt.Println()

	root.postOrder()
	fmt.Println()
}
