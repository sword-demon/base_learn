package main

import "strings"

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	// 如果这是叶子节点，
	// 那么匹配上之后就可以调用该方法
	handler handlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {

}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	// 切割 /
	// user/friends
	// [user, friends]

	// 处理前后的斜杠
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root
	for _, path := range paths {
		mathChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = mathChild
		} else {
			// 没找着, 创建
			h.createSubTree()
		}
	}
}

//func (n *node) findMatchChild(path string) (*node, bool) {
//	// 遍历所有的children
//	for _, child := range n.children {
//		if child.path == path {
//			return child, true
//		}
//	}
//	return nil, false
//}

func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	// 遍历所有的children
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

func (h *HandlerBasedOnTree) createSubTree() {

}