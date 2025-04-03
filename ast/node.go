package ast

// Node is the interface implemented by all LaTeX AST nodes.
// Pos returns the starting byte offset in the original input.
// End returns the byte offset immediately after the node.
type Node interface {
	Pos() int
	End() int
}
