package ast

// Walk traverses an AST starting from the given node
func Walk(v Visitor, node Node) {
	if node == nil {
		return
	}

	v.Visit(node)
	node.VisitChildren(v)
}
