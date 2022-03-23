package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	// 中序遍历 左中右
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
