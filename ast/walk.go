package ast

// Walk traverses an AST starting from the given node,
// calling the visitor for each node.
func Walk(v Visitor, node Node) {
	if node == nil {
		return
	}

	// The Accept method will call the appropriate Visit method on the visitor
	node.Accept(v)
}
