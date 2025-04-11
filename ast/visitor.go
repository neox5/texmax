package ast

// Visitor defines the interface for visiting AST nodes
type Visitor interface {
	// Visit processes a node
	Visit(node Node, role string)

	// EnterNode is called when entering a container node, before visiting children
	EnterNode(node Node)

	// ExitNode is called when exiting a container node, after visiting children
	ExitNode(node Node)
}
