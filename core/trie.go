package core

type node struct {
	pattern  string  //路由全路径
	part     string  //路由的部分值
	children []*node //子节点
	isWild   bool    //是否精准匹配
}

// 第一个匹配成功的节点  用于插入
func (n node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
