package tree

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	// 中序遍历 左中右
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
